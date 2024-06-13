package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	n := screenshot.NumActiveDisplays()
	println(n)
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		fmt.Println(bounds)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
		client := gosseract.NewClient()
		defer client.Close()

		client.SetImage("screenshot.png")
		text, _ := client.Text()
		fmt.Println(text)

	}
}
