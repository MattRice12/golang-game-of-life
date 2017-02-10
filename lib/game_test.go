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
	grid := BuildGrid(&game)
	assert.Equal(t, []int{0, 0, 0, 0}, grid)
}

func TestGeneration(t *testing.T) {
	game := CreateGame(4, 1)
	grid := BuildGrid(&game)
	for i := 0; i < game.Rows; i++ {
		Generation(&grid[i])
	}

	assert.Equal(t, []int{1, 1, 1, 1}, grid)
}
