package main

import (
	"fmt"
	"os"

	"github.com/tkanos/gonfig"
)

var configAssh Configuration

func loadConfigFromDotfile() {
	configuration := Configuration{}
	err := gonfig.GetConf(getConfPath(), &configuration)
	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}
	configAssh = configuration
}

func getConfPath() string {
	user := os.Getenv("USER")

	return "/home/" + user + "/.assh"
}
