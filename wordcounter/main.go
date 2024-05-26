package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("./main.go")
	if err != nil {
		log.Fatalln(err)
	}
	_, kk, err := bufio.ScanBytes(file, true)
	if err != nil {

	}
	for _, i := range kk {
		println(i)
	}

}
