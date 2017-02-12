package main

import "github.com/mattrice12/game-of-life/lib"

func main() {
	width := 40
	height := 20
	patterns := lib.CreatePatterns()
	start := patterns.Pulsar
	game := lib.CreateGame(width, height, start)
	lib.RunGame(game)
}
