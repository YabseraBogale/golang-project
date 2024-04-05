package main

import (
	"fmt"
	"math/rand"
)

func worldviwer(grid *[6][6]int) {
	println("======")
	for i := range grid {
		for j := range grid {
			fmt.Print(grid[i][j])
		}
		println()
	}
	println("======")
}

func main() {
	grid := [6][6]int{}
	buildworld(&grid)
	for i := 0; i < 10; i++ {
		worldviwer(&grid)
		census(&grid)
		worldviwer(&grid)
	}
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
				if sum < 2 || sum == 4 {
					grid[i][j] = 0
					grid[i][j+1] = 0
					grid[i+1][j] = 0
					grid[i+1][j+1] = 0
				} else if sum == 3 {
					if grid[i+1][j+1] == 0 {

					}
				}
			}
		}
	}

}
