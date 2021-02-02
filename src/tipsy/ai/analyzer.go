package ai

import (
	"tipsy/game"
)

const (
	NUMBER_OF_PUCK = 6
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
		if puck.Flipped == true && puck.Color == game.BLUE {
			flippedBluePuck++
		}
		if puck.Color == "black" {
			blackPuck++
		}
	}
	if blackPuck == 0 {
		return currentGame.CurrentPlayer
	}
	if flippedRedPuck == NUMBER_OF_PUCK && flippedBluePuck == NUMBER_OF_PUCK {
		panic("Invalid pucks configuration, all pucks could not be flipped at the same time")
	}
	if flippedRedPuck == NUMBER_OF_PUCK {
		return game.RED
	}
	if flippedBluePuck == NUMBER_OF_PUCK {
		return game.BLUE
	}

	return "active"
}
