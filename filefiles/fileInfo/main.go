package main

import (
	"time"
	"log"
	"os"
)

func main() {
	fileinfo, err := os.Stat("../createFile/log.txt")

	if err != nil {
		log.Fatal(err)
	}
	println("File Name:", fileinfo.Name())
	println("File Size: ", fileinfo.Size())
	println("File Time", fileinfo.ModTime().String())
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err=os.Chtimes("../createFile/log.txt",lastAccessTime,lastModifyTime)
	if err!=nil{
		log.Fatal(err)
	}
}
