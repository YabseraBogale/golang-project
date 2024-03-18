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
	db,_:=sql.Open("sqlite3",data.Name())
	row,_:=db.Query("select * from Software;")
	defer row.Close()
	col,_:=row.Columns()
	for _,i:=range col{
		fmt.Println(i)
	}
}
