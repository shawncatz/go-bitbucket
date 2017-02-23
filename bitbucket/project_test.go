package bitbucket

import (
	"testing"
)

func TestClient_Projects(t *testing.T) {
	list, err := testClient.Projects()
	assertError(t, err)
	assertStringEquals(t, list.Values[0].Key, "chef")
}

func TestClient_Project(t *testing.T) {
	proj, err := testClient.Project("test")
	assertError(t, err)
	assertStringEquals(t, proj.Description, "Description")
}

func TestClient_ProjectMissing(t *testing.T) {
	_, err := testClient.Project("blarg")
	if err == nil {
		t.Errorf("missing project should not exist: %s", err)
	}
}
