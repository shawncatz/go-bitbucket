package bitbucket

import (
	"errors"
	"fmt"

	"github.com/keybase/go-keychain"
	"github.com/tcnksm/go-gitconfig"
)

type Config struct {
	URL      string
	Username string
	Keychain string

	password string
}

func loadConfig(file string) (*Config, error) {
	cfg := &Config{}

	cfg.URL, err = gitconfig.Global("bitbucket.url")
	if err != nil {
		return nil, err
	}

	cfg.Username, err = gitconfig.Global("bitbucket.user")
	if err != nil {
		return nil, err
	}

	cfg.Keychain, err = gitconfig.Global("bitbucket.keychain")
	if err != nil {
		return nil, err
	}

	pwd, err := keychain.GetGenericPassword(cfg.Keychain, cfg.Username, "", "")
	if err != nil {
		return nil, fmt.Errorf("could not get password from keychain: %s", err)
	}
	if string(pwd) == "" {
		return nil, errors.New("password is empty")
	}

	cfg.password = string(pwd)

	return cfg, nil
}
