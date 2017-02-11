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

func TestGenerationRow(t *testing.T) {
	game := CreateGame(4, 1)
	grid := BuildGrid(game)
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			Generation(&grid, x, y)
		}
	}
	assert.Equal(t, [][]int{{0, 0, 0, 1}}, grid)

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			Generation(&grid, x, y)
		}
	}
	assert.Equal(t, [][]int{{0, 0, 1, 0}}, grid)
}

func TestGenerationGrid(t *testing.T) {
	game := CreateGame(3, 3)
	grid := BuildGrid(game)
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			Generation(&grid, x, y)
		}
	}
	assert.Equal(t, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 1}}, grid)

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			Generation(&grid, x, y)
		}
	}
	assert.Equal(t, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 1, 0}}, grid)
}

func TestSetFirstCell(t *testing.T) {
	game := CreateGame(4, 1)
	grid := BuildGrid(game)
	firstCell := SetFirstCell(grid[0])
	assert.Equal(t, 1, firstCell, "All other cells are dead; so firstCell should be alive")
	assert.NotEqual(t, 0, firstCell)
}

func TestStringifyRow(t *testing.T) {
	last := StringifyRow([]int{0, 0, 1})
	middle := StringifyRow([]int{0, 1, 0})
	assert.Equal(t, []string{" ", " ", "@"}, last)
	assert.Equal(t, []string{" ", "@", " "}, middle)
}
