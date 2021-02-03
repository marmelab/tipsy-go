package ai

import (
	"fmt"
	"tipsy/game"
)

const (
	//NumberOfPuck The number of each red and blue pucks
	NumberOfPuck = 6
	ActiveScore  = 0
	WinningScore = 100
	LosingScore  = -100
)

//GetNextMovesScores evaluate each move to find which win or not
func GetNextMovesScores(currentGame game.Game, askingPlayer string) map[string]int {
	directions := [4]string{"right", "left", "up", "down"}
	moves := make(map[string]int)
	board := game.NewBoard()
	for _, firstDirection := range directions {
		fmt.Printf("%v", firstDirection)
		firstMoveGame := game.Tilt(currentGame, &board, firstDirection)
		score := GetScore(firstMoveGame, askingPlayer)
		if score == ActiveScore {
			fmt.Println()
			for _, secondDirection := range directions {
				fmt.Printf("|-- %v", secondDirection)
				secondMoveGame := game.Tilt(firstMoveGame, &board, secondDirection)
				score := GetScore(secondMoveGame, askingPlayer)
				if score != ActiveScore {
					fmt.Printf(" => %v win", score)
					moves[firstDirection+":"+secondDirection] = score
				}
				fmt.Println()
			}
		} else {
			fmt.Printf(" => %v win\n", score)
			moves[firstDirection] = score
		}
	}
	return moves
}

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

	return ActiveScore
}

func getAskingPlayerScore(winningPlayer string, askingPlayer string) int {
	if winningPlayer == askingPlayer {
		return WinningScore
	}
	return LosingScore
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
		if puck.Flipped == true && puck.Color == game.RED {
			flippedRedPuck++
		}
		if puck.Flipped == true && puck.Color == game.BLUE {
			flippedBluePuck++
		}
	}
	return flippedRedPuck, flippedBluePuck
}
