package main

import (
	"fmt"
	"github.com/howeyc/gopass"
)

func prompt(text, def string) (response string) {
	defPrompt := ""
	if def != "" {
		defPrompt = fmt.Sprintf(" [%s]", def)
	}

	fmt.Printf("%s%s: ", text, defPrompt)

	fmt.Scanln(&response)
	if response == "" {
		response = def
	}
	return response
}

func passwd(text string) (response []byte) {
	fmt.Print(text + ": ")
	response, _ = gopass.GetPasswdMasked()
	return response
}
