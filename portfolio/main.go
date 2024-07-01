package main

import (
	"net/http"
)

const (
	port = ":8000"
)

func home(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "./templates/home.html")
}
func articles(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "./templates/articles.html")
}
func contact(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "./templates/contact.html")
}
func project(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "./templates/project.html")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/articles", articles)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/project", project)
	http.ListenAndServe(port, nil)

}
