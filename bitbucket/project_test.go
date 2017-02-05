package bitbucket

import (
	"testing"
)

func TestProjects(t *testing.T) {
	_, err := testClient.Projects()
	assertError(t, err)
}

func TestProject(t *testing.T) {
	_, err := testClient.Project("chef")
	assertError(t, err)
}

func TestProjectMissing(t *testing.T) {
	_, err := testClient.Project("blarg")
	if err == nil {
		t.Errorf("missing project should not exist: %s", err)
	}
}
