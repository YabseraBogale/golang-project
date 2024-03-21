package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	stackmap := make(map[int][]string)
	stack := []string{"django", "flask", "fastapi", "javascript", "php", "wordpress", "java", "spring boot",
		"spring", "node", "front end", "frontend", "back end", "back end", "fullstack", "react", "vue", "c#", ".net", "dotnet",
		"asp.net", "python", "bot", "andriod", "ios", "mobile", "mysql", "mongodb", "postgres", "flutter", "dart", "angularjs"}

	data, err := os.Open("freelance-data-v2.db")
	if err != nil {
		log.Println(err)
	}
	defer data.Close()
	db, _ := sql.Open("sqlite3", data.Name())
	row, _ := db.Query("select message,id from Software;")
	defer row.Close()
	for row.Next() {
		var message string
		var id int
		row.Scan(&message, &id)
		message = strings.ToLower(message)
		for _, i := range stack {
			state := strings.Contains(message, i)
			if state {
				stackmap[id] = append(stackmap[id], i)
			}
		}
	}
	for key, value := range stackmap {
		print(key, " ")
		for _, i := range value {
			print(i, " ")
		}
		println()
	}
}
