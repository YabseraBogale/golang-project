package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("./IRIS.csv")
	if err != nil {
		fmt.Println("openning file")
	}

	data := csv.NewReader(file)
	for {
		d, err := data.Read()
		if err == io.EOF {
			fmt.Println("End of file")
			break
		}
		if err != nil {
			fmt.Println("Err from For loop")
		}
		for i := range d {
			fmt.Println(d[i])
		}
	}

}
