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
func BuildGrid(g Game) [][]int {
	grid := [][]int{}
	for i := 0; i < g.Cols; i++ {
		grid = append(grid, []int{})
		for j := 0; j < g.Rows; j++ {
			grid[i] = append(grid[i], 0)
		}
	}
	return grid
}

// RunGame runs the game
func RunGame(g Game) {
	for {
		grid := BuildGrid(g)
		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid[x]); y++ {
				Generation(&grid, x, y)
				PrintGeneration(&grid)
				time.Sleep(time.Second / 20)
			}
		}
	}
}

// grid = [][]int
// x = int
// y = int
// locate by: grid[x][y]

// Generation changes the ecosystem for each round
func Generation(grid *[][]int, x int, y int) {
	if y == 0 {
		(*grid)[x][y] = SetFirstCell((*grid)[x])
		if x != 0 {
			(*grid)[x-1][len((*grid)[x])-1] = 0
		}
	} else {
		(*grid)[x][y], (*grid)[x][y-1] = (*grid)[x][y-1], (*grid)[x][y]
	}
}

// SetFirstCell determines if first cell is alive or dead
func SetFirstCell(cols []int) int {
	for _, cell := range cols {
		if cell == 1 {
			return 0
		}
	}
	return 1
}

// PrintGeneration prints out what each generation looks like
func PrintGeneration(g *[][]int) {
	clearScreen()
	for i := 0; i < len(*g); i++ {
		gridString := StringifyRow((*g)[i])
		for _, a := range gridString {
			fmt.Print(a)
		}
		fmt.Println()
	}
}

// StringifyRow turns integer into string
func StringifyRow(row []int) []string {
	arr := []string{}
	for _, b := range row {
		if b == 0 {
			arr = append(arr, " ")
		} else {
			arr = append(arr, "@")
		}
	}
	return arr
}

func clearScreen() {
	print("\033[H\033[2J")
}
