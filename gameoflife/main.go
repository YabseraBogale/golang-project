package main

import (
	"fmt"
	"math/rand"
)

func main() {
	grid := [6][6]string{}
	buildworld(&grid)
	// census(&grid)
	fmt.Println(grid[0][1] == " ")
}

func buildworld(grid *[6][6]string) {
	for i := range grid {
		for j := range grid {
			number := rand.Intn(3)
			if number == 0 {
				grid[i][j] = " "
			} else if number == 1 {
				grid[i][j] = "|"
			} else if number == 2 {
				grid[i][j] = "_"
			}
		}
	}
}

// func census(grid *[6][6]int) {
// 	grid[0][1] = 100

// }
