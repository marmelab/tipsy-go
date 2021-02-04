package ai

import (
	"fmt"
	"sync"
	"tipsy/game"
)

//MovementScore tuple of movement with its score
type MovementScore struct {
	movement string
	score    int
}

//GetNextMovesScores evaluate each move to find which win or not
func GetNextMovesScores(currentGame game.Game, depth int, verbose bool) (string, map[string]int) {

	moves := make(map[string]int)
	movesChannel := make(chan MovementScore, 50)
	var wg sync.WaitGroup
	board := game.NewBoard()
	for _, firstDirection := range game.Directions {
		firstMoveGame := game.Tilt(currentGame, &board, firstDirection)
		score := GetScore(firstMoveGame)
		if score != WinningScore && score != LosingScore {
			if verbose {
				fmt.Println()
			}

			for _, secondDirection := range game.Directions {
				wg.Add(1)
				go func(movesChannel chan<- MovementScore, firstDirection, secondDirection string, wg *sync.WaitGroup) {

					secondMoveGame := game.Tilt(firstMoveGame, &board, secondDirection)
					score := MinMax(secondMoveGame, depth, false, false)
					movesChannel <- MovementScore{
						movement: firstDirection + ":" + secondDirection,
						score:    score}
					wg.Done()
				}(movesChannel, firstDirection, secondDirection, &wg)
			}
		} else {
			if verbose {
				fmt.Printf(" => %v\n", score)
			}
			moves[firstDirection] = score
		}
	}
	wg.Wait()
	close(movesChannel)
	bestMove := ""
	bestScore := -9999999
	for moveScore := range movesChannel {
		moves[moveScore.movement] = moveScore.score
		if moveScore.score > bestScore {
			bestScore = moveScore.score
			bestMove = moveScore.movement
		}
		if verbose {
			fmt.Printf(" %v => %v", moveScore.movement, moveScore.score)
			fmt.Println()
		}
	}
	return bestMove, moves
}
