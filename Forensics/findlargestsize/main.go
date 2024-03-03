package main

import (
	"log"
	"os"
)

func main(){
	fileinfo,err:=os.ReadDir("../..")
	if err!=nil{
		log.Fatal(err)
	}
	for _,i:=range fileinfo{
		if !i.IsDir(){
			name:="../../"+i.Name()
			t,err:=os.Stat(name)
			if err!=nil{
				log.Fatal(err)
			} else{
				println(t.Size())
			}

		} else{

		}
	}
}
