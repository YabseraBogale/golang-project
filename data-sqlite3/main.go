package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main(){
	data,err:=os.Open("freelance-data-v2.db")
	if err!=nil{
		log.Println(err)
	}
	defer data.Close()
	db,_:=sql.Open(data)
	row,_:=db.Query("select * from Software;")
	fmt.Println(row)
}
