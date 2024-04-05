package main

import (
	"math/rand"
)

func main() {
	grid := [6][6]int{}
	buildworld(&grid)
	census(&grid)

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
				if sum == 0 {

				} else if sum == 2 || sum == 3 {

				} else if sum == 4 {

				}
			}
		}
	}

}
