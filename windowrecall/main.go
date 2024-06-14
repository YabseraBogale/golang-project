package main

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	n := screenshot.NumActiveDisplays()
	println(n)
	for i := 0; i < time.Now().Day(); i++ {
		bounds := screenshot.GetDisplayBounds(1)
		fmt.Println(bounds)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		timeTaken := time.Now()
		fileName := fmt.Sprintf("%d-%d-%d_%dx%d.png", timeTaken.Day(), timeTaken.Hour(), timeTaken.Minute(), bounds.Dx(), bounds.Dy())
		time.Sleep(30 * 1000)
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)
		fmt.Printf(bounds.String(), fileName)

	}
}
