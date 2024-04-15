package main

import (
	"fmt"
	"sync"

	"github.com/skip2/go-qrcode"
)

func main() {
	urls := []string{
		"www.example.com/1",
		"www.example.com/2",
		"www.example.com/3",
	}
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for i, url := range urls {
		url := url
		i := i
		go func() {
			defer wg.Done()
			generateQR(i, url)
		}()
	}
	wg.Wait()
}

func generateQR(i int, url string) {
	qrCode, _ := qrcode.New(url, qrcode.Medium)
	fileName := fmt.Sprintf("%d.png", i)
	err := qrCode.WriteFile(256, fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("QR code generated and saved as %v.png", i))
}
