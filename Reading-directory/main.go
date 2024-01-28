package main

import (
	"os"
)

func main() {
	file, err := os.Open("./hello")
	if err != nil {
		println("Does n't exist")
	} else {
		println(file.Name())
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
