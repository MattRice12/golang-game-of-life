package main

import "github.com/mattrice12/game-of-life/lib"

func main() {
	game := lib.CreateGame(40, 20)

	lib.RunGame(game)
}
