package bitbucket

import (
	"testing"
)

func TestReposService_List(t *testing.T) {
	list, err := testClient.Repos.List("test", &ReposListOptions{ListOptions: ListOptions{}})
	assertError(t, err)
	assertIntegerEquals(t, list.Values[0].ID, 1)
	assertStringEquals(t, list.Values[0].Name, "My repo")
}

func TestReposService_Get(t *testing.T) {
	repo, err := testClient.Repos.Get("test", "testing")
	assertError(t, err)
	assertIntegerEquals(t, repo.ID, 1)
	assertStringEquals(t, repo.Slug, "my-repo")
	assertStringEquals(t, repo.Name, "My repo")
}

func TestReposService_Create(t *testing.T) {
	repo, err := testClient.Repos.Create("test", "testing", "Testing")
	assertError(t, err)
	assertIntegerEquals(t, repo.ID, 1)
	assertStringEquals(t, repo.Slug, "my-repo")
	assertStringEquals(t, repo.Name, "My repo")
}
