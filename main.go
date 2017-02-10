package main

import "github.com/mattrice12/game-of-life/lib"

func main() {
	game := lib.CreateGame(1, 1)
	lib.RunGame(&game)
}
