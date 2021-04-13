package main

import (
	"covid-19/pkg"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {

	router := pkg.InintService()
	router.Run()
}
