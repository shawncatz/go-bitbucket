package bitbucket

import "testing"

func TestClient_Repositories(t *testing.T) {
	assertNotNil(t, "testClient", testClient)
	_, err := testClient.Repositories("test")
	assertError(t, err)
}

func TestClient_Repository(t *testing.T) {
	assertNotNil(t, "testClient", testClient)
	_, err := testClient.Repository("test", "testing")
	assertError(t, err)
}
