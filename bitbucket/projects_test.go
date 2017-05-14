package bitbucket

import (
	"testing"
)

func TestProjectsService_List(t *testing.T) {
	list, err := testClient.Projects.List()
	assertError(t, err)
	assertStringEquals(t, list.Values[0].Key, "PRJ")
}

func TestProjectsService_Get(t *testing.T) {
	proj, err := testClient.Projects.Get("test")
	assertError(t, err)
	assertStringEquals(t, proj.Description, "The description for my cool project.")
}

func TestProjectsService_GetMissing(t *testing.T) {
	_, err := testClient.Projects.Get("blarg")
	if err == nil {
		t.Errorf("missing project should not exist: %s", err)
	}
}

func TestProjectsService_Create(t *testing.T) {
	proj, err := testClient.Projects.Create("test", "Test", "Description")
	assertError(t, err)
	assertStringEquals(t, proj.Description, "The description for my cool project.")
}

func TestProjectsService_Delete(t *testing.T) {
	err := testClient.Projects.Delete("test")
	assertError(t, err)
}

func TestProjectsService_Update(t *testing.T) {
	old := &Project{Name: "Old", Key: "old"}
	proj, err := testClient.Projects.Update("test", old)
	assertError(t, err)
	assertStringEquals(t, proj.Description, "The description for my cool project.")
}
