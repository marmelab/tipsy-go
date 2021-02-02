package tests

import (
	"testing"
	"tipsy/game"
)

func TestObstaclesShouldNotBePartOfTheBoard(t *testing.T) {
	//GIVEN
	var board = game.NewBoard()

	obstacles := [][]int{
		{0, 3}, {1, 1}, {1, 5}, {2, 2},
		{2, 4}, {3, 0}, {3, 6}, {4, 2},
		{4, 4}, {5, 1}, {5, 5}, {6, 3}}
	//THEN
	for _, obstacle := range obstacles {
		if game.Contains(obstacle, board) == true {
			t.Errorf("Obstacle should not be in the graph %v ", obstacle)
		}
	}
}
