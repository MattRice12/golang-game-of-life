package lib

import (
	"fmt"
	"time"
)

// Game sets up the game board with Rows and Columns that will be manipulated
type Game struct {
	Rows int
	Cols int
}

// CreateGame creates the board
func CreateGame(x int, y int) Game {
	rows := x
	cols := y
	game := Game{rows, cols}
	return game
}

// RunGame runs the game
func RunGame(g *Game) {
	grid := BuildGrid(g)

	for {
		print("\033[H\033[2J")
		for i := 0; i < g.Rows; i++ {
			Generation(&grid[i])
			fmt.Print(grid[i])
			fmt.Println("")
		}
		time.Sleep(time.Second)
	}
}

// BuildGrid builds the grid
func BuildGrid(g *Game) []int {
	grid := []int{}
	for i := 0; i < g.Rows; i++ {
		grid = append(grid, 0)
	}
	return grid
}

// Generation looks at a cell in each generation and changes it according to the rules
func Generation(cell *int) {
	if *cell == 0 {
		*cell = 1
	} else {
		*cell = 0
	}
}
