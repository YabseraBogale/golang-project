package main

import "net"

func main() {
	// go to make use ip resolver
	ip, err := net.ResolveIPAddr("ip", "bbc.com")
	if err != nil {

	}
// how to use ipaddress to find or make my own version of shodan
	println(ip.String())

}
