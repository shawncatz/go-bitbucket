package bitbucket

import (
	"testing"
)

func TestClient_Projects(t *testing.T) {
	_, err := testClient.Projects()
	assertError(t, err)
}

func TestClient_Project(t *testing.T) {
	_, err := testClient.Project("test")
	assertError(t, err)
}

func TestClient_ProjectMissing(t *testing.T) {
	_, err := testClient.Project("blarg")
	if err == nil {
		t.Errorf("missing project should not exist: %s", err)
	}
}
