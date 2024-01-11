package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// ...
func main() {

	
	if err != nil {
		panic(err)
	} else if err = db.Ping(); err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.Exec("create table user( name varchar(20), password varchar(20))")
}
