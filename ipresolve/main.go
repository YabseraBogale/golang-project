package main

import (
	"net"
)

func main() {
	// go to make use ip resolver
	ip, err := net.ResolveIPAddr("ip", "bbc.com")
	if err != nil {

	}
	println("ip: " + ip.String())

}
