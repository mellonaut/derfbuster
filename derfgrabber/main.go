package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "Opened a can...")

// fmt.Println("[!] ERROR")
// fmt.Println("No input parameters provided")
// set global vars from results of previous gobuster,feroxbuster,etc
// ip=
// check_http=
// check_https= 
// check_port=
	
// function for checking input
	// check each host in the scan results for $web_ports or $web_portsplus (including other popular/easily exploitable services known to run on weird ports
	// if $port is found, dial that port to verify host is up
	// create directory for target in the "loot" folder
	// banner grab / timestamp / something to bookmark a kickoff and store

	
