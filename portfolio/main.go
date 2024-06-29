package main

import (
	"fmt"
	"net/http"
)

const (
	port = ":8000"
)

func home(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintln(writer, "home")
}
func articles(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "articles")
}
func contact(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "contact")
}
func project(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "project")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/articles", articles)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/project", project)
	http.ListenAndServe(port, nil)
}
