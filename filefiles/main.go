package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(f)
	f.Close()
}
