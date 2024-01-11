package app

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ...
func main() {

	db, err := sql.Open("mysql", "yabsera:Deathandlife13579@/app")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
