package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello World net")
	resp, err := http.Get("https://stackoverflow.com/questions/41858635/segmentation-violation-with-golang-channels")
	if err != nil {
		fmt.Println("1")
		os.Exit(1)
	}
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("2")
		os.Exit(1)
	}
	for _, data := range b[1:10] {

		fmt.Printf("%b", data)
	}
}
