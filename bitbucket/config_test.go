package bitbucket

import "testing"

func TestNewConfig(t *testing.T) {
	cfg, err := newConfig("fixtures/config.json")
	if err != nil {
		t.Errorf("creating config: %s", err)
	}

	assertStringEquals(t, cfg.Username, "scatanzarite")
}
