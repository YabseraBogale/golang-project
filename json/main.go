package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "data.db")
	defer db.Close()
	row, _ := db.Exec("select * Data where id=" + strconv.Itoa(rand.Intn(4000)))
	fmt.Println(row)
}
