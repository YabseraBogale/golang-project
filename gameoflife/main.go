package main

import (
	"fmt"
	"math/rand"
)

func main() {
	grid := [6][6]string{}
	buildworld(&grid)
	// census(&grid)
	fmt.Println(grid)
}

func buildworld(grid *[6][6]string) {
	for i := range grid {
		for j := range grid {
			if rand.Intn(3) == 0 {
				grid[i][j] = "0"
			} else if rand.Intn(3) == 1 {
				grid[i][j] = "|"
			} else if rand.Intn(3) == 2 {
				grid[i][j] = "__"
			}
		}
	}
}

// func census(grid *[6][6]int) {
// 	grid[0][1] = 100

// }
