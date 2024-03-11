package main

import (
	"fmt"
	"os"
)

func main(){
	fileinfo,err:=os.ReadDir(".")
	if err!=nil{

	} else{
		for _,i:=range fileinfo{
			f,err:=os.Stat(i.Name())
			if err!=nil{

			} else{
				fmt.Printf("Name: %s\nSize: %.2f\n",f.Name(),float32(f.Size()/1000000))
			}
		}
	}

}
