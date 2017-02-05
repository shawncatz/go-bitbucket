package bitbucket

import "testing"

var (
	testClient *Client
	err        error
)

func init() {
	testClient, err = New("fixtures/config.json")
	if err != nil {
		panic("error creating client")
	}
}

func assertError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func assertStringEquals(t *testing.T, is, want string) {
	if is != want {
		t.Errorf("assert string equals: %s != %s", is, want)
	}
}
