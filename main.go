package main

import "github.com/mattrice12/game-of-life/lib"

func main() {
	var (
		start   string
		pattern [][]int
		width   int
		height  int
	)
	pattern, width, height = lib.PromptPattern(&start)
	game := lib.CreateGame(width, height, pattern)
	lib.RunGame(game)
}
