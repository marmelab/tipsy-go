package ai

import "tipsy/game"

const (
	//NeutralScore The neutral score
	NeutralScore = 0
	//WinningScore the score when winning
	WinningScore = 1000
	//LosingScore the score when losing
	LosingScore      = -1000
	flippedPuckScore = 10
	numberOfPuck     = 6
	oneHop           = 100
	twoHops          = 80
	threeHops        = 40
	fourHops         = 20
	fiveHops         = 10
)

var heatMap = map[string]int{
	"0:0": twoHops, "0:1": threeHops, "0:2": threeHops, "0:4": twoHops, "0:5": oneHop, "0:6": twoHops,
	"1:0": oneHop, "1:2": fourHops, "1:3": fourHops, "1:4": threeHops, "1:6": threeHops,
	"2:0": twoHops, "2:1": threeHops, "2:3": fiveHops, "2:5": fourHops, "2:6": threeHops,
	"3:1": fourHops, "3:2": fiveHops, "3:3": fiveHops, "3:4": fiveHops, "3:5": fourHops,
	"4:0": threeHops, "4:1": fourHops, "4:3": fiveHops, "4:5": threeHops, "4:6": twoHops,
	"5:0": threeHops, "5:2": threeHops, "5:3": fourHops, "5:4": fourHops, "5:6": oneHop,
	"6:0": twoHops, "6:1": oneHop, "6:2": twoHops, "6:4": threeHops, "6:5": threeHops, "6:6": twoHops}

//GetScore return the winner of the game, or active if no winner yet
func GetScore(currentGame game.Game, playerAsking string) int {

	fallenRedPucks, fallenBluePucks, fallenBlackPuck := getFallenPucks(currentGame)
	flippedRedPuck, flippedBluePuck := getFlippedPucks(currentGame)

	if fallenBlackPuck {
		return getAskingPlayerScore(currentGame.CurrentPlayer, playerAsking)
	}
	if flippedRedPuck == numberOfPuck && flippedBluePuck == numberOfPuck {
		panic("Invalid pucks configuration, all pucks could not be flipped at the same time")
	}
	if flippedRedPuck == numberOfPuck {
		return getAskingPlayerScore(game.RED, playerAsking)
	}
	if flippedBluePuck == numberOfPuck {
		return getAskingPlayerScore(game.BLUE, playerAsking)
	}
	if currentGame.CurrentPlayer == game.BLUE {
		if flippedRedPuck+fallenRedPucks == numberOfPuck {
			return getAskingPlayerScore(game.RED, playerAsking)
		}
		if flippedBluePuck+fallenBluePucks == numberOfPuck {
			return getAskingPlayerScore(game.BLUE, playerAsking)
		}
	}
	if currentGame.CurrentPlayer == game.RED {
		if flippedBluePuck+fallenBluePucks == numberOfPuck {
			return getAskingPlayerScore(game.BLUE, playerAsking)
		}
		if flippedRedPuck+fallenRedPucks == numberOfPuck {
			return getAskingPlayerScore(game.RED, playerAsking)
		}
	}

	return getActiveScore(currentGame, playerAsking)
}

func getActiveScore(currentGame game.Game, askingPlayer string) int {
	score := NeutralScore
	for key, puck := range currentGame.Pucks {
		if !puck.Flipped {
			switch puck.Color {
			case game.BLACK:
				if currentGame.CurrentPlayer == askingPlayer {
					score -= heatMap[key] * 10
				} else {
					score += heatMap[key] * 10
				}
			case askingPlayer:
				score += heatMap[key]
			default:
				score -= heatMap[key]
			}
		} else {
			switch puck.Color {
			case askingPlayer:
				score += flippedPuckScore
			default:
				score -= flippedPuckScore
			}
		}
	}
	return score
}

func getAskingPlayerScore(winningPlayer string, askingPlayer string) int {
	if winningPlayer == askingPlayer {
		return WinningScore
	}
	return LosingScore
}
