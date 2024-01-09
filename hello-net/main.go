package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello World net")
	resp, err := http.Get("www.google.com")
	if err != nil {
		fmt.Println(err)
	}
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	for _, data := range b[1:] {
		fmt.Println(data)
	}
}
