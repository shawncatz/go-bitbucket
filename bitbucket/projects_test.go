package bitbucket

import (
	"testing"
)

func TestClient_Projects(t *testing.T) {
	list, err := testClient.Projects.List()
	assertError(t, err)
	assertStringEquals(t, list.Values[0].Key, "chef")
}

func TestClient_Project(t *testing.T) {
	proj, err := testClient.Projects.Get("test")
	assertError(t, err)
	assertStringEquals(t, proj.Description, "Description")
}

func TestClient_ProjectMissing(t *testing.T) {
	_, err := testClient.Projects.Get("blarg")
	if err == nil {
		t.Errorf("missing project should not exist: %s", err)
	}
}
