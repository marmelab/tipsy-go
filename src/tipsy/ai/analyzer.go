package ai

import (
	"fmt"
	"tipsy/game"
)

//GetNextMovesScores evaluate each move to find which win or not
func GetNextMovesScores(currentGame game.Game, depth int, verbose bool) map[string]int {
	moves := make(map[string]int)
	board := game.NewBoard()
	for _, firstDirection := range game.Directions {
		if verbose {
			fmt.Printf("%v", firstDirection)
		}
		firstMoveGame := game.Tilt(currentGame, &board, firstDirection)
		score := GetScore(firstMoveGame)
		if score != WinningScore && score != LosingScore {
			if verbose {
				fmt.Println()
			}
			for _, secondDirection := range game.Directions {
				if verbose {
					fmt.Printf("|-- %v", secondDirection)
				}
				secondMoveGame := game.Tilt(firstMoveGame, &board, secondDirection)
				score := MinMax(secondMoveGame, depth, false, false)

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
