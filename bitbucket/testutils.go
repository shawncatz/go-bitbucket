package bitbucket

import (
	"fmt"
	"testing"

	"github.com/shawncatz/go-bitbucket/mock"
)

var (
	testServer *mock.TestBitbucketServer
	testClient *Client
)

func createTestClient(url, user, pass string) *Client {
	testClient = NewClient(url, user, pass)
	testClient.setTesting()

	return testClient
}

func createTestServer() *mock.TestBitbucketServer {
	testServer = mock.NewBitbucket()
	testServer.HandleFile("GET", "/projects", "fixtures/projects.json")
	testServer.HandleFile("POST", "/projects", "fixtures/project.json")
	testServer.HandleFile("GET", "/projects/test", "fixtures/project.json")
	testServer.HandleFile("DELETE", "/projects/test", "fixtures/project.json")
	testServer.HandleFile("PUT", "/projects/test", "fixtures/project.json")
	testServer.HandleFile("GET", "/projects/test/repos", "fixtures/repositories.json")
	testServer.HandleFile("POST", "/projects/test/repos", "fixtures/repository.json")
	testServer.HandleFile("GET", "/projects/test/repos/testing", "fixtures/repository.json")
	testServer.HandleFile("GET", "/projects/test/repos/testing/pull-requests", "fixtures/pull-requests.json")
	testServer.HandleFile("GET", "/inbox/pull-requests", "fixtures/pull-requests.json")
	testServer.HandleFile("GET", "/inbox/pull-requests/count", "fixtures/inbox_count.json")

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

func assertIntegerEquals(t *testing.T, is, want int) {
	if is != want {
		t.Errorf("assert integer equals: %d != %d", is, want)
	}
}
