package bitbucket

import "testing"

func TestInboxService_List(t *testing.T) {
	list, err := testClient.Inbox.List()
	assertError(t, err)
	assertStringEquals(t, list.Values[0].Title, "Talking Nerdy")
}

func TestInboxService_Count(t *testing.T) {
	count, err := testClient.Inbox.Count()
	assertError(t, err)
	assertIntegerEquals(t, count, 1)
}
