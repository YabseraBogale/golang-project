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
	fs, err := os.OpenFile("log.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			println("Permmsion denided")
		}
	}
	fs.Close()
}
