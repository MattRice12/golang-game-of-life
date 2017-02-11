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

// BuildGrid builds the grid with all cells dead
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

// InitialBuild takes a multidimensional array with coordinates for living cells
func InitialBuild(g *[][]int, arr [][]int) [][]int {
	for s := range arr {
		x := arr[s][0]
		y := arr[s][1]
		(*g)[x][y] = 1
	}
	return *g
}

// RunGame runs the game
func RunGame(g Game) {
	grid := BuildGrid(g)
	start := [][]int{{2, 1}, {2, 2}, {2, 3}, {1, 3}, {0, 2}}
	InitialBuild(&grid, start)

	for i := 0; i < 1000; i++ {
		PrintGeneration(grid)
		time.Sleep(time.Second / 20)

		newGrid := BuildGrid(g)
		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid[x]); y++ {
				cell := Generation(grid, x, y)
				newGrid[x][y] = cell
			}
		}
		grid = newGrid
	}
}

// Generation determines whether the cell lives or dies based on the amount of living/dead neighbors
func Generation(g [][]int, x int, y int) int {
	if x > 0 && x < len(g)-1 && y > 0 && y < len(g[x])-1 {
		count := 0
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i != 0 && j != 0 {
					if g[x+i][y+j] == 1 {
						count++
					}
				}
			}
		}
		c := Rules(count)
		return c
	}
	return 0
}

// Rules takes the living neighbor count and determines whether a cell is alive or dead
func Rules(c int) int {
	if c < 1 || c > 3 {
		return 0
	}
	return 1
}

// PrintGeneration prints out what each generation looks like
func PrintGeneration(g [][]int) {
	clearScreen()
	for i := 0; i < len(g); i++ {
		gridString := StringifyRow((g)[i])
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
