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
	census(&grid)
	fmt.Println(grid)

}

func census(grid *[6][6]int) {
	grid[0][1] = 100

}
