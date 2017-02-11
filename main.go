package main

import "github.com/mattrice12/game-of-life/lib"

func main() {
	game := lib.CreateGame(16, 8)
	// grid := lib.BuildGrid(game)
	lib.RunGame(game)
}
