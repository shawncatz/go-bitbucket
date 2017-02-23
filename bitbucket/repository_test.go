package bitbucket

import (
	"testing"
)

func TestClient_Repositories(t *testing.T) {
	list, err := testClient.Repositories("test")
	assertError(t, err)
	assertIntegerEquals(t, list.Values[0].ID, 1234)
	assertStringEquals(t, list.Values[0].Name, "Testing")
}

func TestClient_Repository(t *testing.T) {
	repo, err := testClient.Repository("test", "testing")
	assertError(t, err)
	assertIntegerEquals(t, repo.ID, 1234)
	assertStringEquals(t, repo.Slug, "testing")
	assertStringEquals(t, repo.Name, "Testing")
}
