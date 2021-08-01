package main

import (
	cmd "covid-19/command"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	cmd.InintCommand()
}
