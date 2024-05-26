package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("./main.go")
	if err != nil {
		log.Fatalln(err)
	}
	print(len(file))

}
