package main

import (
	"net"
)

func main() {
	// go to make use ip resolver
	ip, err := net.ResolveIPAddr("ip", "bbc.com")
	if err != nil {

	}
<<<<<<< HEAD
// how to use ipaddress to find or make my own version of shodan
	println(ip.String())
=======
	println("ip: " + ip.String())
>>>>>>> 6e49f01a11065266d77e11a80637328780ec6ddd

}
