package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

)

func main(){
	file,err:=os.Open("fact.json")
	if err!=nil{
		log.Println(err)
	}
	defer file.Close()
	data:=bufio.NewReader(file)
	d,err:=data.Peek(100)
	if err!=nil{
		log.Println(err)
	}
	for _,i:=range d{
		fmt.Print(string(i))
	}
}
