package main

import (
	"net"
)

func main() {
	// now test the csv files
	ip, err := net.Dial("ip4", "192.0.2.1")
	if err != nil {

	}
	println(ip.LocalAddr().String())

}
