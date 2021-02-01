package tests

import (
	"testing"
	"tipsy/game"
)

func TestObstaclesShouldNotBePartOfTheBoard(t *testing.T) {
	//GIVEN
	var board = game.NewBoard()

	obstacles := [][2]int{
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

func TestObstaclesShouldNotHaveEdgesToOthers(t *testing.T) {
	//GIVEN
	obstacles := [][]int{
		{0, 3}, {1, 1}, {1, 5}, {2, 2},
		{2, 4}, {3, 0}, {3, 6}, {4, 2},
		{4, 4}, {5, 1}, {5, 5}, {6, 3}}

	//WHEN
	var board = game.NewBoard()

	//THEN
	if len(board.Edges) == 0 {
		t.Errorf("Edges should be initialized")
	}
	for _, edge := range board.Edges {
		for _, obstacle := range obstacles {
			var to = edge.To.Position
			var from = edge.From.Position
			if (to[0] == obstacle[0] && to[1] == obstacle[1]) || (from[0] == obstacle[0] && from[1] == obstacle[1]) {
				t.Errorf("Obstacle should not be in the graph %v %v", obstacle, edge)
			}
		}
	}
}

func TestExitShouldHaveExitAttribute(t *testing.T) {
	//GIVEN
	exits := [][2]int{
		{1, -1}, {7, 1}, {-1, 5}, {5, 7}}

	//WHEN
	var board = game.NewBoard()

	for _, exit := range exits {
		for _, vertex := range board.Vertices {
			if (vertex.Position[0] == exit[0]) && (vertex.Position[1] == exit[1]) && vertex.Exit != true {
				t.Errorf("Exit should have exit attribute to true %v", exit)
			}
		}
	}
}
