package main

import (
	"github.com/WST-T/Silence/internal/game"
)

func main() {
	// Create and start a new game
	game := game.NewGame()
	game.Play()
}
