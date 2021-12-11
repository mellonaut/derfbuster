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
// derf has a while/do here, but I'm not sure how to do that in Go
// h) help
// v) version
// s) server address
// f) enumeration file
// e) enumeration file type
// P) Polite mode - will show preview and ask for confirmation
// \?) Invalid Options
// If no options given send to help
// if [[ $server_param == "" ]]; then
//     novalues
//     exit 1
	
// set global vars from results of previous gobuster,feroxbuster,etc
// ip=
// check_http=
// check_https= 
// check_port=
	
// function for checking input
	// check each host in the scan results for $web_ports or $web_portsplus (including other popular/easily exploitable services known to run on weird ports
	// if $port is found, dial that port to verify host is up
	// Check if site is HTTP
	// Check if site is HTTPS
	// Check if enumeration types are supported
	  
// Create variables for each enumeration type
func create_vars()
// Generate links for Gobuster Use
// If gobuster format parses, print the link in the file and pass it back to $server_param
// Else, throw error w/ "does not appear to be a gobuster file"
	
// Generate links for Ferox Use
// If ferox format parses, print the link in the file and pass it back to $server_param
// Else, throw error w/ "does not appear to be a ferox file"

// Generate links for FFuf Use
// If ffuf format parses, print the link in the file and pass it back to $server_param
// Else, throw error w/ "does not appear to be a ffuf file"

// Generate links for Dirsearch / Dirb/ Dirbuster
// Generate links for Burp Suite / Zap

// if FFuf / Ferox file will, double-check if links port is same as given port
// Double check actions to be performed
func review(){

	fmt.Println("---------------------------------REVIEW---------------------------------")
	fmt.Println("Input file             $file_param")
	fmt.Println("File type:             $enum_type")
	fmt.Println("Results Dir:           ./$folder_name/")
	fmt.Println("Image Files Dir:       ./$folder_name/$images")
	fmt.Println("Aggregated page:       ./$folder_name/$agg_site")
	fmt.Println("")
	fmt.Println("Sites to grab screen shots from:")
	fmt.Println("$links")
	fmt.Println(-----------------------------------------------------------------------)
	fmt.Println("")
	fmt.Println("[+] Grabbing site screen shots, standby!")
	}
// if in Polite mode, pass the review to stdout and await user conf, err

// If no errors, run the actions
	
// for each host that is up
    
	// if target up create directory for target in the $loot folder 	
    func create_directories() {
		// Make directories to put screen shots in
		if [ $check_http ]; then
			mkdir -p ./$folder_name/http-site-images/
		else
			mkdir -p ./$folder_name/https-site-images/
		fi
	} 
	
	// GO banner grab / timestamp / something to bookmark a kickoff and store
	      
	// start crawling the urls for the target in the scanfile
	func grab_screen_shots_and_aggregate()
	
	// screenshot each page and aggregate using cutycpt?
	        // grab html source for each page
	        // dump screenshots, source dumps
	      // check and signal all tasks done signal to close goroutine
	    // generate html report from results, zip screenshots and raw files with just the banner grab and the report
	      
	// If using Optional attached DB or mounted volume can do something for that here
		  
	// If using Optional attached message broker / queue can do something for that here
		 
	// Pass off to middleware / API Handler module for processing
	

	
