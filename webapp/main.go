package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.ServeFile(http.Dir("./static"))
	http.Handle("/", fs)
	log.Println("listen at 30000")
}
