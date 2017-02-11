package main

import "github.com/mattrice12/game-of-life/lib"

func main() {
	width := 30
	height := 15
	start := [][]int{{3, 2}, {3, 3}, {3, 4}, {2, 4}, {1, 3}}

	game := lib.CreateGame(width, height, start)
	lib.RunGame(game)
}
