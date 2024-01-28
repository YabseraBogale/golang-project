package main

import (
	"os"
	"strings"
)

func main() {
	nn, err := os.Getwd()
	if err != nil {

	} else {
		result := strings.Split(nn, "/")
		println(len(result))
		for _, i := range result {
			if i != "" {
				println(i)
			}

		}
	}
}

func look(name string) bool {

	file, err := os.Open(name)
	if err != nil {
		name = "../" + name
		again := look(name)
		return again
	} else {
		println("File exists" + file.Name())
		return true
	}
}
