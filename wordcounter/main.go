package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.ReadFile("./main.go")
	if err != nil {
		log.Fatalln(err)
	}
	for _, i := range file {
		println(strconv.Itoa(i))
	}

}
