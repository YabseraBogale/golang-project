package main

import (
	"os"
	"log"
)



func main(){
	fileinfo,err:=os.ReadDir("../..")
	if err!=nil{
		log.Fatal(err)
	}
	for _,i:=range fileinfo{
		println(i.Info())
	}

}
