package mock

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// TestBouncerServer is a mock implementation of BouncerClient
type TestBitbucketServer struct {
	mux  *mux.Router
	port int
}

// NewTestBouncerServer returns a mock implementation of a bouncer server
func NewBitbucket() *TestBitbucketServer {
	return &TestBitbucketServer{mux: mux.NewRouter(), port: 8888}
}

// Start starts up the mock bouncer server
func (server *TestBitbucketServer) Start() error {
	var err error = nil
	go func() {
		url := fmt.Sprintf(":%d", server.port)
		err = http.ListenAndServe(url, server.mux)
		if err != nil {
			fmt.Printf("HTTP SERVER: %s\n", err)
		}
	}()

	return err
}

func (server *TestBitbucketServer) URL() string {
	return fmt.Sprintf("http://localhost:%d", server.port)
}

// Port returns the port server is listening on
func (server *TestBitbucketServer) Port() int {
	return server.port
}

// HandleFunc install an endpoint handler
func (server *TestBitbucketServer) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) {
	server.mux.HandleFunc(path, handler)
}

func (server *TestBitbucketServer) HandleFile(method, path, file string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadFile(file)
		if err != nil {
			w.Write([]byte(`{"error": "could not read file"}`))
			return
		}
		w.Write(body)
	}
	server.mux.HandleFunc(path, handler).Methods(method)
}
