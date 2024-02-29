package main

import (
	"log"
	"os"
)

func main() {
	//err := os.Rename("../createFile/log.txt", "../createFile/works.txt")
	//if err != nil {
	///	log.Fatal(err)
	//}
	//returnnig to back
	err := os.Rename("../createFile/works.txt", "../createFile/log.txt")
	if err != nil {
		log.Fatal(err)
	}
}
