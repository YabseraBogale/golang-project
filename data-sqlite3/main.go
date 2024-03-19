package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main(){
	data,err:=os.Open("freelance-data-v2.db")
	if err!=nil{
		log.Println(err)
	}
	defer data.Close()
	db,_:=sql.Open("sqlite3",data.Name())
	row,_:=db.Query("select * from Software;")
	defer row.Close()
	count:=0
	for row.Next(){
		var id int
		var message string
		var date time.Time
		row.Scan(&id,&message,&date)
		fmt.Println(id,message,date)
		count+=1
	}
	fmt.Println("count is",count)

}
