package ai

import (
	"fmt"
	"sync"
	"tipsy/game"
)

const (
	//NeutralScore The neutral score
	NeutralScore = 0
	//WinningScore the score when winning
	WinningScore = 2000
	//LosingScore the score when losing
	LosingScore      = -2000
	blackPuckScore   = 20
	flippedPuckScore = 10
	numberOfPuck     = 6
	oneHop           = 100
	twoHops          = 50
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
func GetScore(currentGame game.Game, remainingTurns bool) int {

	fallenRedPucks, fallenBluePucks, fallenBlackPuck := getFallenPucks(currentGame)
	flippedRedPuck, flippedBluePuck := getFlippedPucks(currentGame)

	if fallenBlackPuck {
		return WinningScore
	}
	if flippedRedPuck == numberOfPuck {
		return getCurrentPlayerScore(game.RED, currentGame.CurrentPlayer)
	}
	if flippedBluePuck == numberOfPuck {
		return getCurrentPlayerScore(game.BLUE, currentGame.CurrentPlayer)
	}
	if currentGame.CurrentPlayer == game.BLUE && !remainingTurns {
		if flippedRedPuck+fallenRedPucks == numberOfPuck {
			return getCurrentPlayerScore(game.RED, currentGame.CurrentPlayer)
		}
		if flippedBluePuck+fallenBluePucks == numberOfPuck {
			return getCurrentPlayerScore(game.BLUE, currentGame.CurrentPlayer)
		}
	}
	if currentGame.CurrentPlayer == game.RED && !remainingTurns {
		if flippedBluePuck+fallenBluePucks == numberOfPuck {
			return getCurrentPlayerScore(game.BLUE, currentGame.CurrentPlayer)
		}
		if flippedRedPuck+fallenRedPucks == numberOfPuck {
			return getCurrentPlayerScore(game.RED, currentGame.CurrentPlayer)
		}
	}

	return getActiveScore(currentGame)
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

func getFallenPucks(currentGame game.Game) (int, int, bool) {

	fallenRedPucks := 0
	fallenBluePucks := 0
	fallenBlackPuck := false
	gameFallenPucks := currentGame.FallenPucks
	for _, puck := range gameFallenPucks {
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

func getActiveScore(currentGame game.Game) int {
	score := NeutralScore
	for key, puck := range currentGame.Pucks {
		if !puck.Flipped {
			switch puck.Color {
			case game.BLACK:
				score -= heatMap[key] * blackPuckScore
			case currentGame.CurrentPlayer:
				score += heatMap[key]
			default:
				score -= heatMap[key]
			}
		}
	}
	return score
}

func getCurrentPlayerScore(winningPlayer string, askingPlayer string) int {
	if winningPlayer == askingPlayer {
		return WinningScore
	}
	return LosingScore
}

//MinMax evaluate best move giving a depth and a starting game
func MinMax(inputGame game.Game, depth int, maximizingPlayer bool, verbose bool) int {
	currentGame := game.CloneGame(inputGame)
	currentGameScore := GetScore(currentGame, false)
	if depth == 0 || currentGameScore == WinningScore || currentGameScore == LosingScore {
		return currentGameScore
	}
	board := game.NewBoard()
	if maximizingPlayer {
		bestScore := -9999999
		var wg sync.WaitGroup
		scoresChannel := make(chan int, 16)
		for _, firstDirection := range game.Directions {
			firstMoveGame := game.ReplacePucks(currentGame)
			firstMoveGame = game.Tilt(currentGame, &board, firstDirection)
			score := GetScore(firstMoveGame, true)
			if score != WinningScore && score != LosingScore {
				for _, secondDirection := range game.Directions {
					wg.Add(1)
					go func(secondDirection string, wg *sync.WaitGroup, scoresChannel chan<- int) {
						defer wg.Done()
						firstMoveGame = game.Tilt(firstMoveGame, &board, secondDirection)
						firstMoveGame = game.SwitchPlayer(currentGame)

						nextGameScore := MinMax(firstMoveGame, depth-1, false, verbose)
						if verbose {
							for i := 0; i < 4-depth; i++ {
								fmt.Print("\t")
							}
							fmt.Printf("Exploring %v %v => %v\n", firstDirection, secondDirection, nextGameScore)
						}
						scoresChannel <- nextGameScore
					}(secondDirection, &wg, scoresChannel)
				}
			} else {
				if score > bestScore {
					bestScore = score
				}
			}
		}
		wg.Wait()
		close(scoresChannel)
		for score := range scoresChannel {
			if score > bestScore {
				bestScore = score
			}

		}
		return bestScore
	}

	var wg sync.WaitGroup
	scoresChannel := make(chan int, 16)
	bestScore := 9999999
	for _, firstDirection := range game.Directions {
		firstMoveGame := game.Tilt(currentGame, &board, firstDirection)
		score := GetScore(firstMoveGame, true)
		if score != WinningScore && score != LosingScore {
			for _, secondDirection := range game.Directions {
				wg.Add(1)
				go func(secondDirection string, wg *sync.WaitGroup, scoresChannel chan<- int) {
					defer wg.Done()
					// nextGame = game.ReplacePucks(nextGame)
					firstMoveGame = game.Tilt(firstMoveGame, &board, secondDirection)
					firstMoveGame = game.SwitchPlayer(currentGame)
					nextGameScore := MinMax(firstMoveGame, depth-1, true, verbose)

					if verbose {
						for i := 0; i < 4-depth; i++ {
							fmt.Print("\t")
						}
						fmt.Printf("Exploring %v %v => %v\n", firstDirection, secondDirection, nextGameScore)
					}
					scoresChannel <- nextGameScore
				}(secondDirection, &wg, scoresChannel)
			}
		} else {
			if score < bestScore {
				bestScore = score
			}
		}
	}
	wg.Wait()
	close(scoresChannel)
	for score := range scoresChannel {
		if score < bestScore {
			bestScore = score
		}
	}
	return bestScore

}
