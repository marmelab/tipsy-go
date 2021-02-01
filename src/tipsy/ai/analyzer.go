package ai

import (
	"tipsy/game"
)

func GameState(currentGame game.Game) string {

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
		return "red"
	}
	if flippedBluePuck == len(currentGame.BluePucks) {
		return "blue"
	}

	return "active"
}
