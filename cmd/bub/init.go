package main

import (
	"fmt"
	"github.com/keybase/go-keychain"
	"github.com/urfave/cli"
	"os/exec"
)

func init() {
	cmdList = append(cmdList, cli.Command{
		Name:        "init",
		Usage:       "initialize configuration",
		Description: `add configuration to gitconfig file and save password to keychain`,
		Action:      cmdInit,
	})
}

func cmdInit(c *cli.Context) error {
	url := prompt("BitBucket URL", "https://api.bitbucket.com/")
	user := prompt("Username", "")
	service := prompt("Keychain Service", "bub")
	pass := passwd("Password")

	fmt.Println("Saving values to git config...")

	_, err := gitConfig("bitbucket.url", url)
	if err != nil {
		return err
	}

	_, err = gitConfig("bitbucket.user", user)
	if err != nil {
		return err
	}

	_, err = gitConfig("bitbucket.keychain", service)
	if err != nil {
		return err
	}

	fmt.Println("Saving password to keychain...")

	item := keychain.NewGenericPassword(service, user, "", pass, "")
	item.SetAccessible(keychain.AccessibleWhenUnlocked)

	_ = keychain.DeleteGenericPasswordItem(service, user)

	err = keychain.AddItem(item)
	if err != nil {
		return err
	}

	return nil
}

func gitConfig(key, value string) ([]byte, error) {
	cmd := []string{"git", "config", "--global", key, value}
	return exec.Command(cmd[0], cmd[1:]...).Output()
}
