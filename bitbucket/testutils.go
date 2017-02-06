package bitbucket

import (
	"fmt"
	"testing"
)

var (
	testClient *Client
	err        error
)

func init() {
	testClient, err = New("fixtures/config.json")
	if err != nil {
		fmt.Printf("error creating client: %s\n", err)
	}
}

func assertNotNil(t *testing.T, name string, val interface{}) {
	if val == nil {
		t.Errorf("%s == nil", name)
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
