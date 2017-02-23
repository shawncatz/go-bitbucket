package bitbucket

import (
	"fmt"
	"testing"

	"github.com/shawncatz/go-bitbucket/mock"
)

var (
	err        error
	testServer *mock.TestBitbucketServer
	testClient *Client
)

func createTestClient() *Client {
	testClient, err = New(&Config{URL: "http://localhost:8888", Username: "user", Password: "password"})
	if err != nil {
		fmt.Printf("error creating client: %s\n", err)
	}

	testClient.SetTesting()

	return testClient
}

func createTestServer() *mock.TestBitbucketServer {
	testServer = mock.NewBitbucket()
	testServer.HandleFile("/projects", "fixtures/projects.json")
	testServer.HandleFile("/projects/test", "fixtures/project.json")
	testServer.HandleFile("/projects/test/repos", "fixtures/repositories.json")
	testServer.HandleFile("/projects/test/repos/testing", "fixtures/repository.json")

	if err := testServer.Start(); err != nil {
		fmt.Printf("could not start server: %s", err)
	}

	return testServer
}

func assertNotNil(t *testing.T, name string, val interface{}) {
	if val == nil {
		t.Errorf("%s == nil", name)
	}
}

func assertError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func assertStringEquals(t *testing.T, is, want string) {
	if is != want {
		t.Errorf("assert string equals: %s != %s", is, want)
	}
}
