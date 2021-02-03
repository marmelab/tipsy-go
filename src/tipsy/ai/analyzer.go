package ai

import (
	"fmt"
	"tipsy/game"
)

const (
	//NUMBER_OF_PUCK The number of each red and blue pucks
	NUMBER_OF_PUCK = 6
)

//GetNextMoves evaluate each move to find which win or not
func GetNextMoves(currentGame game.Game) map[string]bool {
	directions := [4]string{"right", "left", "up", "down"}
	moves := make(map[string]bool)
	board := game.NewBoard()
	for _, firstDirection := range directions {
		fmt.Printf("%v", firstDirection)
		firstMoveGame := game.Tilt(currentGame, &board, firstDirection)
		winner := GetWinner(firstMoveGame)
		if winner == game.ACTIVE {
			fmt.Println()
			for _, secondDirection := range directions {
				fmt.Printf("|-- %v", secondDirection)
				secondMoveGame := game.Tilt(firstMoveGame, &board, secondDirection)
				winner := GetWinner(secondMoveGame)
				if winner != game.ACTIVE {
					fmt.Printf(" => %v win", winner)
					moves[firstDirection+":"+secondDirection] = (winner == secondMoveGame.CurrentPlayer)
				}
				fmt.Println()
			}
		} else {
			fmt.Printf(" => %v win\n", winner)
			moves[firstDirection] = (winner == firstMoveGame.CurrentPlayer)
		}
	}
	return moves
}

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

	return game.ACTIVE
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
