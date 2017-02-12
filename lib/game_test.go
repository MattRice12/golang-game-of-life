package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Grids purposefully made 2 dimension larger (both x and y) to provide for cleaner looping
var (
	Initial4x1 = [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0}}

	Initial4x4 = [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0},
		{1, 0, 1, 0, 0, 0},
		{0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0}}
)

func TestCreateGame(t *testing.T) {
	game := CreateGame(1, 1, [][]int{{}})
	assert.NotNil(t, game)
	assert.Equal(t, Game{3, 3, [][]int{{}}}, game)

	game = CreateGame(20, 20, [][]int{{}})
	assert.Equal(t, Game{22, 22, [][]int{{}}}, game)
}

func TestBuildGrid(t *testing.T) {
	game := CreateGame(4, 1, [][]int{{}})
	grid := BlankGrid(game)

	assert.NotNil(t, grid)
	assert.Equal(t, Initial4x1, grid)
}

func TestGeneration(t *testing.T) {
	game := CreateGame(4, 4, [][]int{{}})
	grid := BlankGrid(game)
	start := [][]int{{2, 1}, {2, 2}, {2, 3}, {1, 3}, {0, 2}}
	InitialCells(&grid, start)
	assert.Equal(t, 1, grid[1][2]) // grid[y][x]

	grid = GenNewGrid(game, grid)
	assert.Equal(t, 0, grid[1][2]) // grid[y][x]

	grid = GenNewGrid(game, grid)
	assert.Equal(t, 1, grid[1][2]) // grid[y][x]
}

func TestInitialCells(t *testing.T) {
	game := CreateGame(4, 4, [][]int{{}})
	grid := BlankGrid(game)
	start := [][]int{{2, 1}, {2, 2}, {2, 3}, {1, 3}, {0, 2}}
	InitialCells(&grid, start)

	assert.Equal(t, Initial4x4, grid)
}

func TestReinitialize(t *testing.T) {
	start := [][]int{{2, 1}, {2, 2}, {2, 3}, {1, 3}, {0, 2}}
	game := CreateGame(4, 4, start)
	grid := BlankGrid(game)
	InitialCells(&grid, start)

	// Initial State before any Generations
	assert.Equal(t, Initial4x4, grid)

	for i := 0; i <= 10; i++ {
		grid = GenNewGrid(game, grid)
		InitializeGrid(&grid, game, i)
	}
	// After 10 Generations, grid should be different
	assert.NotEqual(t, Initial4x4, grid)

	// At 20 Generations, grid starting position reinitializes
	InitializeGrid(&grid, game, 20)
	assert.Equal(t, Initial4x4, grid)
}

func TestRules(t *testing.T) {
	start := [][]int{{1, 1}}
	game := CreateGame(4, 4, start)
	grid := BlankGrid(game)
	InitialCells(&grid, start)
	var output int
	// For currently dead cells
	y := 0
	x := 0
	for i := 0; i < 3; i++ {
		output = Rules(i, grid, y, x)
		assert.Equal(t, 0, output, "Expect 0 (death) -- Count == %v.", i)
	}
	output = Rules(3, grid, y, x)
	assert.Equal(t, 1, output, "Expect 1 (life). Count == %v. Rule: 'Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction'", 3)

	// For currently living cells
	y = 1
	x = 1
	for i := 0; i < 2; i++ {
		output = Rules(i, grid, y, x)
		assert.Equal(t, 0, output, "Expect 0 (death). Count == %v. Rule: 'Any live cell with fewer than two live neighbors die, as if by isolation'", i)
	}

	for i := 2; i < 4; i++ {
		output = Rules(i, grid, y, x)
		assert.Equal(t, 1, output, "Expect 1 (life). Count == %v. Rule: 'Any live cell with two or three live neighbors lives on to the next generation'.", i)
	}

	for i := 4; i <= 8; i++ {
		output = Rules(i, grid, y, x)
		assert.Equal(t, 0, output, "Expect 0 (death). Count > 3. Rule: 'Any live cell with more than three live neighbors dies, as if by overcrowding'")
	}
}

func TestStringifyRow(t *testing.T) {
	last := StringifyRow([]int{0, 0, 1})
	middle := StringifyRow([]int{0, 1, 0})
	assert.Equal(t, []string{" ", " ", "*"}, last)
	assert.Equal(t, []string{" ", "*", " "}, middle)
}

func TestFindPattern(t *testing.T) {
	pattern, width, height := findPattern("glider")
	assert.Equal(t, []int{3, 2}, pattern[0])
	assert.Equal(t, 30, width)
	assert.Equal(t, 15, height)

	pattern, width, height = findPattern("pulsar")
	assert.Equal(t, []int{12, 16}, pattern[0])
	assert.Equal(t, 40, width)
	assert.Equal(t, 20, height)

	pattern, width, height = findPattern("gun")
	assert.Equal(t, []int{3, 13}, pattern[0])
	assert.Equal(t, 50, width)
	assert.Equal(t, 30, height)
}
