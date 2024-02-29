package main

import (
	"log"
	"os"
)

func main() {
	fileinfo, err := os.Stat("../createFile/log.txt")

	if err != nil {
		log.Fatal(err)
	}
	println("File Name:", fileinfo.Name())
	println("File Size: ", fileinfo.Size())
	println("File Time", fileinfo.ModTime())

}
