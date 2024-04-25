package main

import (
	"net"
)

func main() {
	// go to make use ip resolver
	ip, err := net.Dial("ip4", "192.0.2.1")
	if err != nil {

	}
	println(ip.LocalAddr().String())

}
