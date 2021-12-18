package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"bufio"
	"flag"	
	"sync"
	"strings"
)

// set global vars we will populate from results of previous gobuster,feroxbuster,etc
// 	BaseURL string IPv4 string port string protocol string
var url string
var ip string
var port string
var protocol string
var current_version string = "0.0.1"
var title string = "Derfbuster"


// Function for no input values	
// func novalues() {
// 	fmt.Println("[!] ERROR")
// 	fmt.Println("No input parameters provided, big dawg. Try harder.")
// }
// func noinput() {
// 	fmt.Print("Enter text: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	// ReadString will block until the delimiter is entered
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("An error occured while reading input. Please try again", err)
// 		return

func crash() {
	f, err := os.Open("/crashlog.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "Opened a can...")

// Version function
func version() {
    fmt.Println(current_version)
}
// Title function
func display_title() {
	fmt.Println(title)
|


// Exampoles
func examples() {
	fmt.Println("Examples:")
	fmt.Println("derfbuster -s http://10.10.10.10 -f <gobusterfile> -e gb")
	fmt.Println("derfbuster -s https://10.10.10.10 -f <ffuf> -e ff")
	fmt.Println("derfbuster -s http://10.10.10.10:8080 -f <feroxbusterfile> -e fb")
	fmt.Println("derfbuster -s https://10.10.10.10:4443 -f <ffuf csv file> -e ff")
}


// Help
func help(){
    version
    fmt.Println("")    
	fmt.Println("Param:      Description:")
    fmt.Println("------      ------------")
    fmt.Println("-h          help")
    fmt.Println("-v          version")
    fmt.Println("-s[svr]     http(s) server to grab screen shots from")
    fmt.Println("-f[file]    Output file from which ever enumeration tool was used")
    fmt.Println("-e[value]   Input file enumeration type. Supported types:")
    fmt.Println("                - Feroxbuster - \"fb\"")
    fmt.Println("                - Gobuster: \"gb\"")
    fmt.Println("                - Ffuf: \"ff\"")
    fmt.Println("            Supported ffuf output formats: json, csv, or md ONLY")
    fmt.Println("            NOTE: you must have jq installed for json formated files to parse")
    examples
}


// function for taking input options arguments etc
func options() {
	h := flag.Bool("h", false, "Get some help.")
	V := flag.Bool("V", false, "Get the current version.")
	v := flag.Bool("v", false, "Verbose mode.")
	s := flag.String("s", "", "HTTP(S) server to grab screen shots from.")
	f := flag.String("f", "", "Output file from which ever enumeration tool was used.")
	e := flag.String("e", "", "Input file enumeration type. Supported types:")
	nr := flag.Bool("nr", false, "Suppress report opening in browser.")
	p := flag.Bool("p", false, "Polite mode enabled")
	vp := flag.Bool("vp", false, "Very Polite mode enabled")
	\? := flag.Bool("\?", false, "Get some help.")
}

// If no options given send to help
// if [[ $server_param == "" ]]; then
//     novalues
//     exit 1
	flag.Parse()
	args := flag.Args()






// ip=
// check_http=
func check_http() {
	BaseURL := ""

	resp, err := http.Get(BaseURL)

	if err != nil {
		fmt.Println(err)
	}

	finalURL := resp.Request.URL.String()
	fmt.Println("Original URL is : ", originalURL)
	fmt.Println("Final URL is : ", finalURL)

	// Check if served with https 
	fmt.Println("Is HTTPS ? : ", strings.HasPrefix(finalURL,"https"))
}


// check_https= 
// check_port= Try starting with this function
func raw_connect(host string, ports []string) {
    for _, port := range ports {
        timeout := time.Second
        conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
        if err != nil {
            fmt.Println("Connecting error:", err)
        }
        if conn != nil {
            defer conn.Close()
            fmt.Println("Opened", net.JoinHostPort(host, port))
        }
    }
}
	
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
		// if [ $check_http ]; then
		// 	mkdir -p ./$folder_name/http-site-images/
		// else
		// 	mkdir -p ./$folder_name/https-site-images/
		// fi
		// So I think here I need to make it
		// if check_http == true  
		check_http, err := os.Mkdir("./$folder_name/http-site-images/")
		if err != nil {
			os.Mkdir("./$folder_name/https-site-images/")
		}
	} 
	
	// GO banner grab / timestamp / something to bookmark a kickoff and store
	      
	// start crawling the urls for the target in the scanfile
	func grab_screen_shots_and_aggregate()
	
	// screenshot each page and aggregate using cutycpt?
	        // grab html source for each page
	        // dump screenshots, source dumps
	      // check and signal all tasks done signal to close goroutine
	    // generate html report from results, zip screenshots and raw files with just the banner grab and the report

	func report() {
			f, err := os.Open("/$folder_name/aggregated.html")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(f.Name(), "Opening report in browser...")

	// If using Optional attached DB or mounted volume can do something for that here
		  
	// If using Optional attached message broker / queue can do something for that here
		 
	// Pass off to middleware / API Handler module for processing
	
// Main Function 
// How do I do Polite mode? Run main but without report? How do I rate limit and do review?
func main() {
	// Check for input parameters, validate
	check_input $server_param $file_param $enum_type
	// Set global vars from input parameters
	create_vars $server_param $file_param $enum_type
	// fmt.Println(title + version)
	title
	
	// Display review / Verify Actions if Polite mode
	review $server_param $file_param $enum_type

	// create target directories in $loot folder
	create_directories 
	// Grab banner grab and store in $loot/$folder_name/
	banner_grab
	// Grab screenshots and store in $loot/$folder_name/$agg_site
	grab_screen_shots_and_aggregate $server_param

	// Pass off to middleware / API Handler module for processing
	hand_off

	fmt.Println("")
	fmt.Println("[+] Aggregated Report saved to ./$folder_name/$agg_site")

	// Open report in browser
	report
}


