package main

import (
	"fmt"
	"math/rand"
)

func worldviwer(grid *[6][6]int) {
	for i := range grid {
		for j := range grid {
			fmt.Print(grid[i][j])
		}
		println()
	}
}

func main() {
	grid := [6][6]int{}
	buildworld(&grid)
	worldviwer(&grid)
	census(&grid)
	worldviwer(&grid)
}

func buildworld(grid *[6][6]int) {
	for i := range grid {
		for j := range grid {
			if rand.Intn(2) == 0 {
				grid[i][j] = 0
			} else {
				grid[i][j] = 1
			}
		}
	}
}

func census(grid *[6][6]int) {
	for i := range grid {
		for j := range grid {
			if i+1 < 6 && j+1 < 6 {
				sum := grid[i][j] + grid[i][j+1] + grid[i+1][j] + grid[i+1][j+1]
				if sum < 2 {
					grid[i][j] = 0
					grid[i][j+1] = 0
					grid[i+1][j] = 0
					grid[i+1][j+1] = 0
				} else if sum == 2 || sum == 3 {

				} else if sum == 4 {
					grid[i][j] = 0
					grid[i][j+1] = 0
					grid[i+1][j] = 0
					grid[i+1][j+1] = 0
				}
			}
		}
	}

}
