package main

import (
	"math"
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
				println(f.Name(),float32(f.Size()/1000000))
			}
		}
	}
	num:=[]float64{1,23,5,-5,-7,8}
	println(math.Max(num[0],num[2]))
}
