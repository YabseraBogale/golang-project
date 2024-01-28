package main

import (
	"os"
)

func main() {
	file, err := os.Open("./hello")
	if err != nil {
		println("Does n't exist")
	}
	println(file.Name())

}
