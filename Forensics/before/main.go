package main

import (
	"log"
	"os"
	"time"
)
func main(){
	name:="/home/yabsera/Documents/Github/golang-project/filefiles/createFile"
	dirInfo,err:=os.ReadDir(name)
	tt:=[]time.Time{}
	if err!=nil{
		log.Fatal(err)
	} else{
		for _,i:=range dirInfo{
			f,err:=os.Stat(name+"/"+i.Name())
			if err!=nil{
				log.Fatalln(err)
			} else{
				println(f.Name())
				tt =append(tt,f.ModTime())
			}
		}
	}
	println(tt[0].Day(),"is before",tt[1].Day(),tt[0].Before(tt[1]))
	println(tt[0].Day(),"is After",tt[1].Day(),tt[0].After(tt[1]))
}
