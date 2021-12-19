package struct_webserver


// struct for webservers being analyzed

type IPv4 struct {
  IP string
}

type URL struct {
  URL string
}

type Port struct {
  Port int
}

type Protocol string {
  Protocol string
}

type webServer struct {
	URL *URL
	IPv4 *IPv4
	Port *Port
	Proto *Protocol
}

type webServer interface {
  scan() // return array of open and closed ports
  connect() bool
  
}
