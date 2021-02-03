package ai

import (
	"tipsy/game"
)

const (
	NUMBER_OF_PUCK = 6
)

//GetWinner return the winner of the game, or active if no winner yet
func GetWinner(currentGame game.Game) string {

	fallenRedPucks, fallenBluePucks, fallenBlackPuck := getFallenPucks(currentGame)
	flippedRedPuck, flippedBluePuck := getFlippedPucks(currentGame)

	if fallenBlackPuck {
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
	if currentGame.CurrentPlayer == game.BLUE {
		if flippedRedPuck+fallenRedPucks == NUMBER_OF_PUCK {
			return game.RED
		}
		if flippedBluePuck+fallenBluePucks == NUMBER_OF_PUCK {
			return game.BLUE
		}
	}
	if currentGame.CurrentPlayer == game.RED {
		if flippedBluePuck+fallenBluePucks == NUMBER_OF_PUCK {
			return game.BLUE
		}
		if flippedRedPuck+fallenRedPucks == NUMBER_OF_PUCK {
			return game.RED
		}
	}

	return "active"
}

func getFallenPucks(currentGame game.Game) (int, int, bool) {

	fallenRedPucks := 0
	fallenBluePucks := 0
	fallenBlackPuck := false
	for _, puck := range currentGame.FallenPucks {
		switch puck.Color {
		case game.BLUE:
			fallenBluePucks++
		case game.RED:
			fallenRedPucks++
		case game.BLACK:
			fallenBlackPuck = true
		}
	}
	return fallenRedPucks, fallenBluePucks, fallenBlackPuck
}
func getFlippedPucks(currentGame game.Game) (int, int) {
	flippedRedPuck := 0
	flippedBluePuck := 0

	for _, puck := range currentGame.Pucks {
		if puck.Flipped == true && puck.Color == "red" {
			flippedRedPuck++
		}
		if puck.Flipped == true && puck.Color == game.BLUE {
			flippedBluePuck++
		}
	}
	return flippedRedPuck, flippedBluePuck
}
