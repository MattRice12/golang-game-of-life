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
	assert.Equal(t, []int{0, 0, 0, 0}, grid)
}

func TestGeneration(t *testing.T) {
	game := CreateGame(4, 1)
	grid := BuildGrid(game)
	for i := 0; i < game.Rows; i++ {
		Generation(grid, i)
	}
	assert.Equal(t, []int{0, 0, 0, 1}, grid)

	for i := 0; i < game.Rows; i++ {
		Generation(grid, i)
	}
	assert.Equal(t, []int{0, 0, 1, 0}, grid)
}

func TestSetFirstCell(t *testing.T) {
	game := CreateGame(4, 1)
	grid := BuildGrid(game)
	firstCell := SetFirstCell(grid)

	assert.Equal(t, 1, firstCell, "All other cells are dead; so firstCell should be alive")

	for i := 0; i < game.Rows; i++ {
		Generation(grid, i)
	}
	firstCell = SetFirstCell(grid)
	assert.Equal(t, 0, firstCell, "After the first generation, firstCell dies")

	cellNeighbor := grid[len(grid)-1]
	assert.Equal(t, 1, cellNeighbor, "In the second generation, firstCell's neighbor (cellNeighbor) comes to life")
}

func TestStringifyCell(t *testing.T) {
	dead := StringifyCell(0)
	alive := StringifyCell(1)
	assert.Equal(t, " ", dead)
	assert.Equal(t, "*", alive)
}
