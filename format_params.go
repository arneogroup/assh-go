package main

import (
	"fmt"
	"os"
	"strings"
)

// Checks if a jumpServer meta exists, otherwise tries to determine the jump server from customer
// Fallback to a default entry is made if none of the above works
func getJumpServer(jumpServer string, hostingOwner string) string {
	var ok bool
	var returnValue string

	if jumpServer != "" {
		returnValue = jumpServer
	} else if returnValue, ok = configAssh.JumpServer[hostingOwner]; !ok {
		returnValue = configAssh.JumpServer["default"]
	}
	return returnValue
}

//Priorise the use of privateIP and fallback to address if no privateIP
func getConnectionIP(privateIP string, address string) string {
	var connectionAddress string

	if connectionAddress = privateIP; connectionAddress == "" {
		connectionAddress = address
	}
	return connectionAddress
}

// If login exists returns it, otherwise use ubuntu as default
func getConnectionUser(login string) string {
	var user string

	if user = login; user == "" {
		user = configAssh.DefaultUsername
	}
	return user
}

// Generate arguments for the rsync command, by replacing the string "_srv_"
// with the valid connection information, adding the jump server if required
func formatRsyncCommandParameters(idxs []int, selectedEntries []ServerInfos) {
	user := getConnectionUser(selectedEntries[0].Meta.Login)
	connectionAddress := getConnectionIP(selectedEntries[0].Meta.PrivateIP, selectedEntries[0].Address)
	jumpServer := getJumpServer(selectedEntries[0].Meta.JumpServer, selectedEntries[0].Meta.HostingOwner)

	fmt.Printf(" -e 'ssh -J %s' %s %s", jumpServer, strings.ReplaceAll(os.Args[1], "_srv_", user+"@"+connectionAddress), strings.ReplaceAll(os.Args[2], "_srv_", user+"@"+connectionAddress))
}

// Generate arguments for the ssh command
func formatSSHCommandParameters(idxs []int, selectedEntries []ServerInfos) {
	for _, idx := range idxs {

		user := getConnectionUser(selectedEntries[idx].Meta.Login)
		connectionAddress := getConnectionIP(selectedEntries[idx].Meta.PrivateIP, selectedEntries[idx].Address)
		jumpServer := getJumpServer(selectedEntries[idx].Meta.JumpServer, selectedEntries[idx].Meta.HostingOwner)

		fmt.Printf("%s@%s -J %s", user, connectionAddress, jumpServer)
	}
}
