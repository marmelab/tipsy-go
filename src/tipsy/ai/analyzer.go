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

	var wg sync.WaitGroup
	moves := make(map[string]int)
	movesChannel := make(chan MovementScore, 16)
	board := game.NewBoard()
	bestMove := ""
	bestScore := -9999999
	for _, firstDirection := range game.Directions {
		firstMoveGame := game.Tilt(currentGame, &board, firstDirection)
		score := GetScore(firstMoveGame, true)
		if score != WinningScore && score != LosingScore {
			for _, secondDirection := range game.Directions {
				wg.Add(1)
				go func(movesChannel chan<- MovementScore, firstDirection, secondDirection string, firstMoveGame game.Game, wg *sync.WaitGroup) {
					fmt.Printf("exploring %v\n", firstDirection+":"+secondDirection)
					secondMoveGame := game.Tilt(firstMoveGame, &board, secondDirection)
					var moveScore int = MinMax(secondMoveGame, depth, false, false)
					movesChannel <- MovementScore{
						movement: firstDirection + ":" + secondDirection,
						score:    moveScore}
					wg.Done()
				}(movesChannel, firstDirection, secondDirection, firstMoveGame, &wg)
			}
			wg.Wait()
		} else {
			moves[firstDirection] = score
			if score > bestScore {
				bestScore = score
				bestMove = firstDirection
			}
		}
	}
	close(movesChannel)
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
