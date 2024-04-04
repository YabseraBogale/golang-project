package main

import "math/rand"

func main() {
	for i := 0; i < 2; i++ {
		println(rand.Intn(2))
	}
}
