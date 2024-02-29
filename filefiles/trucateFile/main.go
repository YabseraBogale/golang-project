package main

import (
	"log"
	"os"
)

func main() {
	// the method take the file name and reduce the file with number of byte decreasing the size and the data in it
	err := os.Truncate("../createFile/log.txt", 20)
	if err != nil {
		log.Fatal(err)
	}

}
