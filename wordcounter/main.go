package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("./main.go")
	if err != nil {
		log.Fatalln(err)
	}
	b := bytes.NewBuffer(file)
	scanner := bufio.NewScanner(b)
	scanner.Split(bufio.ScanBytes)

}
