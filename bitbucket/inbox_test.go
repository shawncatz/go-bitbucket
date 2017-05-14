package bitbucket

import "testing"

func TestInboxService_List(t *testing.T) {
	list, err := testClient.PullRequests.List("test", "testing")
	assertError(t, err)
	assertStringEquals(t, list.Values[0].Title, "Talking Nerdy")
}
