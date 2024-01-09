package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World net")
	resp, err := http.Get("www.google.com")
	if err == nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Body)
}
