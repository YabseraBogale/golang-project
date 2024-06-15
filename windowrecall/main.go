package main

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/pato/screenshot"
)

func main() {
here:
	c1 := make(chan string, 1)
	timeout := 3 * time.Second
	go func() {
		xconn, err := screenshot.Setup()
		if err != nil {
			panic(err)
		}
		img, err := screenshot.CaptureScreen(xconn)
		if err != nil {
			panic(err)
		}
		filename := "./" + string(time.Now().Day()) + string(time.Now().Hour()) + ".png"
		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		err = png.Encode(f, img)
		if err != nil {
			panic(err)
		}
		f.Close()
		time.Sleep(timeout)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		goto here
	}
}
