package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./IRIS.csv")
	if err != nil {

	}
	data := csv.NewReader(file)
	df, err := data.ReadAll()
	if err != nil {
		os.Exit(1)
	}
	for i := range df {
		fmt.Println(df[i])
	}

}
