package main

import (
	"fmt"
	"math/rand"
)

func main() {
	grid := [6][6]int{}
	for i := range grid {
		for j := range grid {
			grid[i][j] = rand.Intn(2)
		}
	}
	fmt.Println(grid)
}
