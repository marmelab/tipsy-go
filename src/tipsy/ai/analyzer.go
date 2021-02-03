package ai

import (
	"fmt"
	"tipsy/game"
)

//GetNextMovesScores evaluate each move to find which win or not
func GetNextMovesScores(currentGame game.Game, verbose bool) map[string]int {
	directions := [4]string{"right", "left", "up", "down"}
	moves := make(map[string]int)
	board := game.NewBoard()
	for _, firstDirection := range directions {
		if verbose {
			fmt.Printf("%v", firstDirection)
		}
		firstMoveGame := game.Tilt(currentGame, &board, firstDirection)
		score := GetScore(firstMoveGame, currentGame.CurrentPlayer)
		if score != WinningScore && score != LosingScore {
			if verbose {
				fmt.Println()
			}
			for _, secondDirection := range directions {
				if verbose {
					fmt.Printf("|-- %v", secondDirection)
				}
				secondMoveGame := game.Tilt(firstMoveGame, &board, secondDirection)
				score := GetScore(secondMoveGame, currentGame.CurrentPlayer)

				if verbose {
					fmt.Printf(" => %v", score)
				}
				moves[firstDirection+":"+secondDirection] = score

				if verbose {
					fmt.Println()
				}
			}
		} else {
			if verbose {
				fmt.Printf(" => %v\n", score)
			}
			moves[firstDirection] = score
		}
	}
	return moves
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
