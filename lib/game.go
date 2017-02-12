package lib

import (
	"fmt"
	"time"
)

// Game sets up the game board with Rows and Columns that will be manipulated
type Game struct {
	Rows  int
	Cols  int
	Start [][]int
}

// CreateGame creates the Game and Grid
func CreateGame(x int, y int, s [][]int) Game {
	rows := x + 2
	cols := y + 2
	game := Game{rows, cols, s}
	return game
}

// BlankGrid builds a grid template with all cells dead
func BlankGrid(g Game) [][]int {
	grid := [][]int{}
	for i := 0; i < g.Cols; i++ {
		grid = append(grid, []int{})
		for j := 0; j < g.Rows; j++ {
			grid[i] = append(grid[i], 0)
		}
	}
	return grid
}

// InitialCells takes a multidimensional array with coordinates for living cells
func InitialCells(grid *[][]int, start [][]int) {
	for i := range start {
		x := &start[i][0]
		y := &start[i][1]
		(*grid)[*y][*x] = 1
	}
}

// Reinitialize calls the initial build every 20 generations
func Reinitialize(grid *[][]int, game Game, i int) {
	if i%20 == 0 {
		InitialCells(grid, game.Start)
	}
}

// RunGame runs the game
func RunGame(game Game) {
	grid := BlankGrid(game)
	for i := 0; i < 1000; i++ {
		Reinitialize(&grid, game, i)
		PrintGeneration(grid)
		grid = GenNewGrid(game, grid)
	}
}

// GenNewGrid makes a blank grid template and populates it
func GenNewGrid(g Game, grid [][]int) [][]int {
	newGrid := BlankGrid(g)
	for y := range grid {
		for x := range grid[y] {
			newGrid[y][x] = Generation(grid, y, x)
		}
	}
	return newGrid
}

// Generation determines whether the cell lives or dies based on the amount of living/dead neighbors
func Generation(grid [][]int, y int, x int) int {
	y, x = BorderCells(grid, y, x)
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				if grid[y+i][x+j] == 1 {
					count++
				}
			}
		}
	}
	cell := Rules(count, grid, y, x)
	return cell
}

// BorderCells allows Pattern to loop once reaching the bottom of the screen
func BorderCells(grid [][]int, y int, x int) (int, int) {
	if y == 0 {
		y = len(grid) - 2
	} else if y == len(grid)-1 {
		y = 1
	}
	if x == 0 {
		x = len(grid[y]) - 2
	} else if x == len(grid[y])-1 {
		x = 1
	}
	return y, x
}

// Rules takes the living neighbor count and determines whether a cell is alive or dead
func Rules(c int, grid [][]int, y int, x int) int {
	if grid[y][x] == 1 {
		if c < 2 || c > 3 {
			return 0
		} else if c == 2 || c == 3 {
			return 1
		}
	}
	if c == 3 {
		return 1
	}
	return 0
}

// PrintGeneration prints out what each generation looks like
func PrintGeneration(grid [][]int) {
	clearScreen()
	for i := 0; i < len(grid)-2; i++ {
		gridString := StringifyRow((grid)[i])
		for i := 0; i < len(gridString)-2; i++ {
			fmt.Print(gridString[i])
		}
		fmt.Println()
	}
	time.Sleep(time.Second / 10)
}

// StringifyRow turns integer into string
func StringifyRow(row []int) []string {
	strArr := []string{}
	for _, cell := range row {
		if cell == 0 {
			strArr = append(strArr, " ")
		} else {
			strArr = append(strArr, "*")
		}
	}
	return strArr
}

func clearScreen() {
	print("\033[H\033[2J")
}
