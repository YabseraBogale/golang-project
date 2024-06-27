package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	log.Println("listen at 30000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println(err.Error())
	}
}
