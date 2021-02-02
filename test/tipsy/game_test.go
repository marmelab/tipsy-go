package tests

import (
	"testing"
	"tipsy/game"
)

func TestGameShouldBeWellDesiaralized(t *testing.T) {
	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#|b| | |#| | | |#",
		"#| |#| | | |#| | ",
		"#| | |#|R|#| | |#",
		"#|#| | |x| | |#|#",
		"#| | |#| |#| | |#",
		"#| |#| | | |#| |#",
		"#| | | |#| | | |#",
		"#################"}
	//WHEN
	currentGame := game.Deserialize(rawGame)
	//THEN
	if currentGame.Pucks[2].Position[0] != 3 || currentGame.Pucks[2].Position[1] != 3 {
		t.Errorf("BlackPuck should be on (3,3) %v", currentGame)
	}
	if currentGame.Pucks[1].Position[0] != 3 || currentGame.Pucks[1].Position[1] != 2 || currentGame.Pucks[1].Color != "red" {
		t.Errorf("RedPuck should be on (3,2) %v", currentGame)
	}
	if currentGame.Pucks[1].Flipped != true {
		t.Errorf("RedPuck should be flipped %v", currentGame)
	}
	if currentGame.Pucks[0].Position[0] != 0 || currentGame.Pucks[0].Position[1] != 0 || currentGame.Pucks[0].Color != "blue" {
		t.Errorf("BluePuck should be on (0,0) %v", currentGame)
	}
	if currentGame.Pucks[0].Flipped != false {
		t.Errorf("BluePuck should not be flipped %v", currentGame)
	}
}
func TestThePuckShouldGoToRightWhenTiltedToEast(t *testing.T) {
	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#|b| | |#| | | |#",
		"#| |#| | | |#| | ",
		"#| | |#| |#| | |#",
		"#|#| | | | | |#|#",
		"#| | |#| |#| | |#",
		"#| |#| | | |#| |#",
		"#| | | |#| | | |#",
		"#################"}
	currentGame := game.Deserialize(rawGame)
	board := game.NewBoard()
	//WHEN
	currentGame = game.Tilt(currentGame, &board, "east")
	//THEN
	if currentGame.Pucks[0].Position[0] != 2 || currentGame.Pucks[0].Position[1] != 0 {
		t.Errorf("Puck should be on (2,0) %v", currentGame)
	}
}
