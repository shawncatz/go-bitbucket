package main

import (
	"github.com/keybase/go-keychain"
	"github.com/shawncatz/go-gitconfig/gitconfig"
	"os"
)

type Config struct {
	URL      string
	User     string
	Keychain string
	password string
}

func loadConfig() (*Config, error) {
	g, err := gitconfig.Read(os.Getenv("HOME") + "/.gitconfig")
	if err != nil {
		return nil, err
	}

	c := &Config{
		URL:      g["bitbucket.url"],
		User:     g["bitbucket.user"],
		Keychain: g["bitbucket.keychain"],
	}

	data, err := keychain.GetGenericPassword(c.Keychain, c.User, "bub", "")
	if err != nil {
		return nil, err
	}

	c.password = string(data)

	return c, nil
}
