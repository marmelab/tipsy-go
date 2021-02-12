package ai

import (
	"context"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for _, firstDirection := range game.Directions {
		firstMoveGame := game.Tilt(currentGame, &board, firstDirection)
		score := GetScore(firstMoveGame, true)
		if score != WinningScore && score != LosingScore {
			for _, secondDirection := range game.Directions {

				secondMoveGame := game.Tilt(firstMoveGame, &board, secondDirection)
				score = GetScore(firstMoveGame, true)
				if score == WinningScore || score == LosingScore {
					moves[firstDirection+":"+secondDirection] = score
					if score > bestScore {
						bestScore = score
						bestMove = firstDirection
						cancel()
					}
				} else {
					wg.Add(1)
					go func(movesChannel chan<- MovementScore,
						firstDirection, secondDirection string,
						secondMoveGame game.Game, wg *sync.WaitGroup,
						ctx context.Context, cancel context.CancelFunc) {
						defer wg.Done()
						maxMoveScore := MinMax(ctx, secondMoveGame, depth, false, false, cancel)
						movesChannel <- MovementScore{
							movement: firstDirection + ":" + secondDirection,
							score:    maxMoveScore}
					}(movesChannel, firstDirection, secondDirection, secondMoveGame, &wg, ctx, cancel)
				}
			}
		} else {
			moves[firstDirection] = score
			if score > bestScore {
				bestScore = score
				bestMove = firstDirection
				cancel()
			}
		}
	}
	wg.Wait()
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
