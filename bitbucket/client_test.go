package bitbucket

import "testing"

func init() {
	testServer = createTestServer()
	testClient = createTestClient()
}

func TestClient(t *testing.T) {
	assertNotNil(t, "testClient", testClient)
}
