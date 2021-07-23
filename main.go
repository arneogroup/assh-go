package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	loadConfigFromDotfile()
	consulList := getListFromConsul()
	selectedEntries := getSelectionFromUser(consulList)
	if len(args) == 0 {
		formatSSHCommandParameters(selectedEntries, consulList)
	} else if len(args) == 2 {
		formatRsyncCommandParameters(selectedEntries, consulList)
	} else {
		log.Fatal("For ssh completion, please provide no arguments. For rsync completion please provide two arguments (use _srv_)")
	}
}
