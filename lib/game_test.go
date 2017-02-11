package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGame(t *testing.T) {
	game := CreateGame(1, 1)
	assert.NotNil(t, game)
	assert.Equal(t, Game{1, 1}, game)

	game = CreateGame(20, 20)
	assert.Equal(t, Game{20, 20}, game)
}

func TestBuildGrid(t *testing.T) {
	game := CreateGame(4, 1)
	grid := BuildGrid(game)
	assert.Equal(t, [][]int{{0, 0, 0, 0}}, grid)
}

func TestGenerationCell(t *testing.T) {
	game := CreateGame(4, 4)
	grid := BuildGrid(game)
	start := [][]int{{2, 1}, {2, 2}, {2, 3}, {1, 3}, {0, 2}}
	InitialBuild(&grid, start)

	cell := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			cell = Generation(grid, x, y)
		}
	}
	assert.Equal(t, 0, cell)

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			cell = Generation(grid, x, y)
		}
	}
	assert.Equal(t, 0, cell)
}

func TestGenerationGrid(t *testing.T) {
	game := CreateGame(4, 4)
	grid := BuildGrid(game)
	start := [][]int{{2, 1}, {2, 2}, {2, 3}, {1, 3}, {0, 2}}
	InitialBuild(&grid, start)
	newGrid := BuildGrid(game)

	assert.Equal(t, [][]int{
		{0, 0, 1, 0},
		{0, 0, 0, 1},
		{0, 1, 1, 1},
		{0, 0, 0, 0}}, grid)

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			cell := Generation(grid, x, y)
			newGrid[x][y] = cell
		}
	}
	assert.Equal(t, [][]int{
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0}}, newGrid)

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			cell := Generation(newGrid, x, y)
			newGrid[x][y] = cell
		}
	}
	assert.Equal(t, [][]int{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0}}, newGrid)
}

func TestStringifyRow(t *testing.T) {
	last := StringifyRow([]int{0, 0, 1})
	middle := StringifyRow([]int{0, 1, 0})
	assert.Equal(t, []string{" ", " ", "@"}, last)
	assert.Equal(t, []string{" ", "@", " "}, middle)
}
