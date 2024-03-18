package main

import (
<<<<<<< HEAD

=======
	"bufio"
	"fmt"
>>>>>>> 3b8d679 (ok)
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

<<<<<<< HEAD

=======
	fmt.Println(data)
>>>>>>> 3b8d679 (ok)
}
