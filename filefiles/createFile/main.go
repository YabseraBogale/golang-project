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
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File does exist. File information:")
	log.Println(fileInfo)
}
