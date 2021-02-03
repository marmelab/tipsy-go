package ai

import (
	"fmt"
	"tipsy/game"
)

//GetNextMovesScores evaluate each move to find which win or not
func GetNextMovesScores(currentGame game.Game, verbose bool) map[string]int {
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
				score := GetScore(secondMoveGame)

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

//MinMax evaluate best move giving a depth and a starting game
func MinMax(currentGame game.Game, depth int, maximizingPlayer bool, verbose bool) (int, string) {
	if depth == 0 {
		return GetScore(currentGame), "None"
	}
	board := game.NewBoard()
	if maximizingPlayer {
		value := -9999999
		directions := "None"
		for _, firstDirection := range game.Directions {
			for _, secondDirection := range game.Directions {
				if verbose {
					fmt.Printf("Exploring %v %v\n", firstDirection, secondDirection)
				}
				nextGame := game.Tilt(currentGame, &board, firstDirection)
				nextGame = game.Tilt(nextGame, &board, secondDirection)
				// nextGame = game.SwitchPlayer(nextGame)
				// nextGame = game.ReplacePucks(nextGame)
				nextGameScore, _ := MinMax(nextGame, depth-1, false, verbose)
				if nextGameScore > value {
					value = nextGameScore
					directions = firstDirection + ":" + secondDirection
				}
			}
		}
		return value, directions
	}

	value := 9999999
	directions := "None"
	for _, firstDirection := range game.Directions {
		for _, secondDirection := range game.Directions {
			nextGame := game.Tilt(currentGame, &board, firstDirection)
			nextGame = game.Tilt(nextGame, &board, secondDirection)
			// nextGame = game.SwitchPlayer(nextGame)
			// nextGame = game.ReplacePucks(nextGame)
			nextGameScore, _ := MinMax(nextGame, depth-1, true, verbose)

			if nextGameScore < value {
				value = nextGameScore
				directions = firstDirection + ":" + secondDirection
			}
		}
	}
	return value, directions

}
