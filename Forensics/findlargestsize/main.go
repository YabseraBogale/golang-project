package main

import "os"

func main(){
	fileinfo,err:=os.ReadDir(".")
	if err!=nil{

	} else{
		for _,i:=range fileinfo{
			f,err:=os.Stat(i.Name())
			if err!=nil{

			} else{
				println(f.Name(),float32(f.Size()/1000000))
			}
		}
	}
}
