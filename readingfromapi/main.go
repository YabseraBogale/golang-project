package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	url := "https://numbersapi.p.rapidapi.com/1492/year?json=true&fragment=true"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "09aedc540emshd7c8ba6f1ecee0ap17ecedjsna99bac53f084")
	req.Header.Add("X-RapidAPI-Host", "numbersapi.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
