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
	wc := 0
	for _, i := range file {
		if string(i) != " " {
			wc += 1
		} else if string(i) == "\n" {
			println("found")
		}
	}
	println(wc)
}
