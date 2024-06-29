package main

import "net/http"

const (
	port = ":8000"
)

func home(writer http.ResponseWriter, request *http.Request) {

}
func articles(writer http.ResponseWriter, request *http.Request) {

}
func contact(writer http.ResponseWriter, request *http.Request) {

}
func project(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/", articles)
	http.HandleFunc("/", contact)
	http.HandleFunc("/", project)
	http.ListenAndServe(port, nil)
}
