package ai

import "tipsy/game"

const (
	//NumberOfPuck The number of each red and blue pucks
	NumberOfPuck = 6
	//NeutralScore The neutral score
	NeutralScore = 0
	//WinningScore the score when winning
	WinningScore = 100
	//LosingScore the score when losing
	LosingScore = -100
)

//GetScore return the winner of the game, or active if no winner yet
func GetScore(currentGame game.Game, playerAsking string) int {

	fallenRedPucks, fallenBluePucks, fallenBlackPuck := getFallenPucks(currentGame)
	flippedRedPuck, flippedBluePuck := getFlippedPucks(currentGame)

	if fallenBlackPuck {
		return getAskingPlayerScore(currentGame.CurrentPlayer, playerAsking)
	}
	if flippedRedPuck == NumberOfPuck && flippedBluePuck == NumberOfPuck {
		panic("Invalid pucks configuration, all pucks could not be flipped at the same time")
	}
	if flippedRedPuck == NumberOfPuck {
		return getAskingPlayerScore(game.RED, playerAsking)
	}
	if flippedBluePuck == NumberOfPuck {
		return getAskingPlayerScore(game.BLUE, playerAsking)
	}
	if currentGame.CurrentPlayer == game.BLUE {
		if flippedRedPuck+fallenRedPucks == NumberOfPuck {
			return getAskingPlayerScore(game.RED, playerAsking)
		}
		if flippedBluePuck+fallenBluePucks == NumberOfPuck {
			return getAskingPlayerScore(game.BLUE, playerAsking)
		}
	}
	if currentGame.CurrentPlayer == game.RED {
		if flippedBluePuck+fallenBluePucks == NumberOfPuck {
			return getAskingPlayerScore(game.BLUE, playerAsking)
		}
		if flippedRedPuck+fallenRedPucks == NumberOfPuck {
			return getAskingPlayerScore(game.RED, playerAsking)
		}
	}

	return NeutralScore
}

func getAskingPlayerScore(winningPlayer string, askingPlayer string) int {
	if winningPlayer == askingPlayer {
		return WinningScore
	}
	return LosingScore
}
