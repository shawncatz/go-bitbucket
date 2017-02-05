package bitbucket

import (
	"encoding/json"
	"fmt"
	"github.com/keybase/go-keychain"
	"io/ioutil"
)

type Config struct {
	URL      string
	Username string
	Keychain string

	password string
}

func newConfig(file string) (*Config, error) {
	cfg := &Config{}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %s: %s", file, err)
	}

	err = json.Unmarshal(data, cfg)
	if err != nil {
		return nil, fmt.Errorf("could not parse json: %s", err)
	}

	pwd, err := keychain.GetGenericPassword(cfg.Keychain, cfg.Username, "", "")
	if err != nil {
		return nil, fmt.Errorf("could not get password from keychain: %s", err)
	}

	cfg.password = string(pwd)

	return cfg, nil
}
