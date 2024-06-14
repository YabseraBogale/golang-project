package main

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {

	bounds := screenshot.GetDisplayBounds(1)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	timeTaken := time.Now()
	fileName := fmt.Sprintf("%d_%dx%d.png", timeTaken.Minute(), bounds.Dx(), bounds.Dy())
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)
	fmt.Println(bounds.String(), fileName)

}
