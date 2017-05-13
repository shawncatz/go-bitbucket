package bitbucket

import (
	"testing"
)

func init() {
	testClient = createTestClient()
	testServer = createTestServer()
}

func TestClient(t *testing.T) {
	assertNotNil(t, "testClient", testClient)
}
