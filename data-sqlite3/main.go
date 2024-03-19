package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main(){
	data,err:=os.Open("freelance-data-v2.db")
	if err!=nil{
		log.Println(err)
	}
	defer data.Close()
	db,_:=sql.Open("sqlite3",data.Name())
	row,_:=db.Query("select message from Software;")
	defer row.Close()
	for row.Next(){
		var message string
		row.Scan(&message)
		state:=strings.Contains(message,"nodejs")
		println(state)
	}


}
