package ai

import (
	"tipsy/game"
)

//GetWinner return the winner of the game, or active if no winner yet
func GetWinner(currentGame game.Game) string {

	if (game.Puck{}) == currentGame.BlackPuck {
		return currentGame.CurrentPlayer
	}

	flippedRedPuck := 0
	flippedBluePuck := 0
	for i := 0; i < len(currentGame.RedPucks); i++ {
		if currentGame.RedPucks[i].Flipped == true {
			flippedRedPuck++
		}
	}
	for i := 0; i < len(currentGame.BluePucks); i++ {
		if currentGame.BluePucks[i].Flipped == true {
			flippedBluePuck++
		}
	}

	if flippedRedPuck == len(currentGame.RedPucks) && flippedBluePuck == len(currentGame.BluePucks) {
		panic("Invalid pucks configuration, all pucks could not be flipped at the same time")
	}
	if flippedRedPuck == len(currentGame.RedPucks) {
		return game.RED
	}
	if flippedBluePuck == len(currentGame.BluePucks) {
		return game.BLUE
	}

	return "active"
}
