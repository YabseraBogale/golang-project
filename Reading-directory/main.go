package main

import "os"

func main() {
	file, err := os.Open("./main.go")
	if err != nil {

	}
	println(file.Name())
}
