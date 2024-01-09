package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello World net")
	resp, err := http.Get("gopl.io/ch1/fetch")
	if err != nil {
		os.Exit(1)
	}
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		os.Exit(1)
	}
	for _, data := range b[1:] {
		fmt.Println(data)
	}
}
