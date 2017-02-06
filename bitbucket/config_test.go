package bitbucket

import "testing"

func TestNewConfig(t *testing.T) {
	cfg, err := loadConfig("fixtures/config.json")
	assertError(t, err)
	assertNotNil(t, "cfg", cfg)
	assertStringEquals(t, cfg.Username, "scatanzarite")
}
