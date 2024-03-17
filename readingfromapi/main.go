package main

import (
	"log"
	"os"
)

func main(){
	file,err:=os.Open("fact.json")
	if err!=nil{
		log.Println(err)
	}
	defer file.Close()

}
