package ai

import (
	"tipsy/game"
)

//GetWinner return the winner of the game, or active if no winner yet
func GetWinner(currentGame game.Game) string {

	flippedRedPuck := 0
	flippedBluePuck := 0
	blackPuck := 0

	for _, puck := range currentGame.Pucks {
		if puck.Flipped == true && puck.Color == "red" {
			flippedRedPuck++
		}
		if puck.Flipped == true && puck.Color == "blue" {
			flippedBluePuck++
		}
		if puck.Color == "black" {
			blackPuck++
		}
	}
	if blackPuck == 0 {
		return currentGame.CurrentPlayer
	}
	if flippedRedPuck == 6 && flippedBluePuck == 6 {
		panic("Invalid pucks configuration, all pucks could not be flipped at the same time")
	}
	if flippedRedPuck == 6 {
		return game.RED
	}
	if flippedBluePuck == 6 {
		return game.BLUE
	}

	return "active"
}
