package bitbucket

import (
	"os"
	"testing"
)

func init() {
	url := os.Getenv("BITBUCKET_URL")
	user := os.Getenv("BITBUCKET_USER")
	pass := os.Getenv("BITBUCKET_PASS")

	if url == "" {
		url = "http://localhost:8888"
		user = "user"
		pass = "password"
		testServer = createTestServer()
	}

	testClient = createTestClient(url, user, pass)
}

func TestClient(t *testing.T) {
	assertNotNil(t, "testClient", testClient)
}
