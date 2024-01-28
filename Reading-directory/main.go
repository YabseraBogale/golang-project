package main

import (
	"os"
	"path"
	"strings"
)

func main() {
	name := path.Dir(".")
	println("name is" + name)
	nn := strings.Split(name, "/")
	println(nn)
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
