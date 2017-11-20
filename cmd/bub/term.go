package main

import (
	"fmt"

	"github.com/agtorre/gocolorize"
	"github.com/howeyc/gopass"
)

type Term struct {
	f map[string]func(...interface{}) string
}

func NewTerm() *Term {
	return &Term{
		f: map[string]func(...interface{}) string{
			"norm": gocolorize.NewColor("white").Paint,
			"bold": gocolorize.NewColor("white+b").Paint,
			"succ": gocolorize.NewColor("green+b").Paint,
			"info": gocolorize.NewColor("blue+b").Paint,
			"warn": gocolorize.NewColor("yellow+b").Paint,
			"crit": gocolorize.NewColor("red+b").Paint,
		},
	}
}

func (t *Term) prompt(text, def string) (response string) {
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

func (t *Term) passwd(text string) (response []byte) {
	fmt.Print(text + ": ")
	response, _ = gopass.GetPasswdMasked()
	return response
}

func (t *Term) Print(color, text string) {
	f := t.f[color]
	fmt.Print(f(text))
}

func (t *Term) Printf(color, format string, a ...interface{}) {
	f := t.f[color]
	fmt.Printf(f(fmt.Sprintf(format, a...)))
}

func (t *Term) Normal(text string) {
	t.Print("norm", text)
}

func (t *Term) Normalf(format string, a ...interface{}) {
	t.Printf("norm", format, a...)
}

func (t *Term) Bold(text string) {
	t.Print("bold", text)
}

func (t *Term) Boldf(format string, a ...interface{}) {
	t.Printf("bold", format, a...)
}

func (t *Term) Success(text string) {
	t.Print("succ", text)
}

func (t *Term) Successf(format string, a ...interface{}) {
	t.Printf("succ", format, a...)
}

func (t *Term) Info(text string) {
	t.Print("info", text)
}

func (t *Term) Infof(format string, a ...interface{}) {
	t.Printf("info", format, a...)
}

func (t *Term) Warn(text string) {
	t.Print("warn", text)
}

func (t *Term) Warnf(format string, a ...interface{}) {
	t.Printf("warn", format, a...)
}

func (t *Term) Error(text string) {
	t.Print("crit", text)
}

func (t *Term) Errorf(format string, a ...interface{}) {
	t.Printf("crit", format, a...)
}
