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


// function for no input values	
// fmt.Println("[!] ERROR")
// fmt.Println("No input parameters provided")
	
// Help, Options, examples, versions functions
	
func examples() {
	fmt.Println("Examples:")
	fmt.Println("\"$ derfgrabber -s http://10.10.10.10 -f <gobusterfile> -e gb\"")
	fmt.Println("\"$ derfgrabber -s https://10.10.10.10 -f <gobusterfile> -e gb\"")
	fmt.Println("\"$ derfgrabber -s http://10.10.10.10:8080 -f <feroxbusterfile> -e fb\"")
	fmt.Println("\"$ derfgrabber -s https://10.10.10.10:4443 -f <ffuf csv file> -e ff\"")
}

	
	
// function for taking input options arguments etc
	
// set global vars from results of previous gobuster,feroxbuster,etc
// ip=
// check_http=
// check_https= 
// check_port=
	
// function for checking input
	// check each host in the scan results for $web_ports or $web_portsplus (including other popular/easily exploitable services known to run on weird ports
	// if $port is found, dial that port to verify host is up
	// if target up create directory for target in the "loot" folder
	  // for each host that is up, 
	    // GO banner grab / timestamp / something to bookmark a kickoff and store
	      // start crawling the urls for the target in the scanfile
	        // screenshot each page
	        // grab html source for each page
	        // dump screenshots, source dumps
	      // check and signal all tasks done signal to close goroutine
	    // generate html report from results, zip screenshots and raw files with just the banner grab and the report
	      // If using Optional attached DB, API or mounted volume can do something for that here, maybe a messaging Q
	

	
