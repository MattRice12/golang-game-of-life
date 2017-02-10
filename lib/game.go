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

// CreateGame creates the Game and Grid
func CreateGame(x int, y int) Game {
	rows := x
	cols := y
	game := Game{rows, cols}
	return game
}

// BuildGrid builds the grid
func BuildGrid(g Game) []int {
	grid := []int{}
	for i := 0; i < g.Rows; i++ {
		grid = append(grid, 0)
	}
	return grid
}

// RunGame runs the game
func RunGame(game Game, grid []int) {
	for {
		for i := 0; i < game.Rows; i++ {
			print("\033[H\033[2J")
			Generation(grid, i)
			PrintGeneration(grid)
		}
		fmt.Println("")
		time.Sleep(time.Second / 10)
	}
}

// Generation changes the ecosystem for each round
func Generation(g []int, i int) {
	if i == 0 {
		g[i] = SetFirstCell(g)
	} else {
		g[i], g[i-1] = g[i-1], g[i]
	}
}

// SetFirstCell determines if first cell is alive or dead
func SetFirstCell(g []int) int {
	for _, cell := range g {
		if cell == 1 {
			return 0
		}
	}
	return 1
}

// PrintGeneration prints out what each generation looks like
func PrintGeneration(g []int) {
	for i := 0; i < len(g); i++ {
		fmt.Print(StringifyCell(g[i]))
	}
}

// StringifyCell turns integer into string
func StringifyCell(c int) string {
	if c == 0 {
		return " "
	}
	return "*"
}
