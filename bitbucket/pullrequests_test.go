package bitbucket

import (
	"testing"
)

func TestPullRequestsService_List(t *testing.T) {
	list, err := testClient.PullRequests.List("test", "testing")
	assertError(t, err)
	assertStringEquals(t, list.Values[0].Title, "Talking Nerdy")
}
