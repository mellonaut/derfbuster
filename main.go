package main

import (
	"fmt"
	"os"
	"net/HTTP"
	"io/ioutil"
	"bufio"
	"flag"	
	"sync"
	"strings"
)

var (
	currentVersion string = "0.0.1"
	title string = "Derfbuster"
)

// set global vars we will populate from results of previous gobuster,feroxbuster,etc
// 	BaseURL string IPv4 string port string protocol string
var (
	URL string
	IPv4 string
	Port string
	Protocol string
)
	


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
    fmt.Println(currentVersion)
}
// Title function
func display_title() {
	fmt.Println(title)
|


// Exampoles
func examples() {
	fmt.Println("Examples:")
	fmt.Println("derfbuster -s HTTP://10.10.10.10 -f <gobusterfile> -e gb")
	fmt.Println("derfbuster -s HTTPS://10.10.10.10 -f <ffuf> -e ff")
	fmt.Println("derfbuster -s HTTP://10.10.10.10:8080 -f <feroxbusterfile> -e fb")
	fmt.Println("derfbuster -s HTTPS://10.10.10.10:4443 -f <ffuf csv file> -e ff")
}


// Help
func help(){
    version
    fmt.Println("")    
	fmt.Println("Param:      Description:")
    fmt.Println("------      ------------")
    fmt.Println("-h          help")
    fmt.Println("-v          version")
    fmt.Println("-s[svr]     HTTP(s) server to grab screen shots from")
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
// if [[ serverParam == "" ]]; then
//     novalues
//     exit 1
	flag.Parse()
	args := flag.Args()


// checkHTTP=
func checkHTTP() {
	BaseURL := ""

	resp, err := HTTP.Get(BaseURL)

	if err != nil {
		fmt.Println(err)
	}

	finalURL := resp.Request.URL.String()
	fmt.Println("Original URL is : ", originalURL)
	fmt.Println("Final URL is : ", finalURL)

	// Check if served with HTTPS 
	fmt.Println("Is HTTPS ? : ", strings.HasPrefix(finalURL,"HTTPS"))
}


// checkHTTPS= 
// checkPort= Try starting with this function
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
func createLinks()
// Generate links for Gobuster Use
// If gobuster format parses, print the link in the file and pass it back to serverParam
// Else, throw error w/ "does not appear to be a gobuster file"
	
// Generate links for Ferox Use
// If ferox format parses, print the link in the file and pass it back to serverParam
// Else, throw error w/ "does not appear to be a ferox file"

// Generate links for FFuf Use
// If ffuf format parses, print the link in the file and pass it back to serverParam
// Else, throw error w/ "does not appear to be a ffuf file"

// Generate links for Dirsearch / Dirb/ Dirbuster
// Generate links for Burp Suite / Zap

// if FFuf / Ferox file will, double-check if links port is same as given port
// Double check actions to be performed
func review(){

	fmt.Println("---------------------------------REVIEW---------------------------------")
	fmt.Println("Input file             fileParam")
	fmt.Println("File type:             enumType")
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
    func createDirs() {
		// Make dirs to put screen shots in
		// if [ checkHTTP ]; then
		// 	mkdir -p ./$folder_name/HTTP-site-images/
		// else
		// 	mkdir -p ./$folder_name/HTTPS-site-images/
		// fi
		// So I think here I need to make it
		// if checkHTTP == true  
		checkHTTP, err := os.Mkdir("./$folder_name/HTTP-site-images/")
		if err != nil {
			os.Mkdir("./$folder_name/HTTPS-site-images/")
		}
	} 
	
	// GO banner grab / timestamp / something to bookmark a kickoff and store
	      
	// start crawling the URLs for the target in the scanfile
	func grabScreens()
	
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
	check_input serverParam fileParam enumType
	// Set global vars from input parameters
	create_vars serverParam fileParam enumType
	// fmt.Println(title + version)
	title
	
	// Display review / Verify Actions if Polite mode
	review serverParam fileParam enumType

	// create target dirs in $loot folder
	createDirs 
	// Grab banner grab and store in $loot/$folder_name/
	bannerGrab
	// Grab screenshots and store in $loot/$folder_name/$agg_site
	grabScreens serverParam

	// Pass off to middleware / API Handler module for processing
	handOff

	fmt.Println("")
	fmt.Println("[+] Aggregated Report saved to ./$folder_name/$agg_site")

	// Open report in browser
	report
}


