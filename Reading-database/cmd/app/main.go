package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// ...
func main() {

	db, err := sql.Open("mysql", "yabsera:Deathandlife13579@/app")
	if err != nil {
		panic(err)
	} else if err = db.Ping(); err != nil {
		panic(err)
	}
	// See "Important settings" section.

	db.Exec("CREATE TABLE IF NOT EXISTS mytable (id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY, some_text TEXT NOT NULL)")
}
