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
	blackPuck, blackPuckExist := currentGame.Pucks["3:3"]
	if !blackPuckExist {
		t.Errorf("BlackPuck should be on (3,3) %v", currentGame)
	}
	if blackPuck.Color != "black" {
		t.Errorf("BlackPuck should be black %v", currentGame)
	}
	redPuck, redPuckExist := currentGame.Pucks["3:2"]
	if !redPuckExist {
		t.Errorf("RedPuck should be on (3,2) %v", currentGame)
	}
	if redPuck.Color != "red" {
		t.Errorf("RedPuck should be red %v", currentGame)
	}
	if redPuck.Flipped != true {
		t.Errorf("RedPuck should be flipped %v", currentGame)
	}
	bluePuck, bluePuckExist := currentGame.Pucks["0:0"]
	if !bluePuckExist {
		t.Errorf("BluePuck should be on (0,0) %v", currentGame)
	}
	if bluePuck.Color != game.BLUE {
		t.Errorf("BluePuck should be blue %v", currentGame)
	}
	if bluePuck.Flipped != false {
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
	bluePuck, bluePuckExist := currentGame.Pucks["2:0"]
	if !bluePuckExist || bluePuck.Color != game.BLUE {
		t.Errorf("BluePuck should be on (2,0) %v", currentGame)
	}
}

func TestThePuckShoulBeStoppedByAnotherPuck(t *testing.T) {
	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#|b| |b|#| | | |#",
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
	bluePuck20, bluePuck20Exists := currentGame.Pucks["2:0"]
	if !bluePuck20Exists || bluePuck20.Color != game.BLUE {
		t.Errorf("Puck should be on (2,0) %v", currentGame)
	}
	bluePuck10, bluePuck10Exists := currentGame.Pucks["1:0"]
	if !bluePuck10Exists || bluePuck10.Color != game.BLUE {
		t.Errorf("Puck should be on (2,0) %v", currentGame)
	}
}

func TestBothPucksShouldMoveToTheSouthAndStoppedByTheWall(t *testing.T) {
	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#| |b| |#| | ",
		"#| | |#|b|#| | |#",
		"#|#| | | | | |#|#",
		"#| | |#| |#| | |#",
		"#| |#| | | |#| |#",
		"#| | | |#| | | |#",
		"#################"}
	currentGame := game.Deserialize(rawGame)
	board := game.NewBoard()

	//WHEN
	currentGame = game.Tilt(currentGame, &board, "south")

	//THEN
	bluePuck34, bluePuck34Exists := currentGame.Pucks["3:4"]
	if !bluePuck34Exists || bluePuck34.Color != game.BLUE {
		t.Errorf("Puck should be on (3,4) %v", currentGame)
	}
	bluePuck35, bluePuck35Exists := currentGame.Pucks["3:5"]
	if !bluePuck35Exists || bluePuck35.Color != game.BLUE {
		t.Errorf("Puck should be on (3,5) %v", currentGame)
	}
}

func TestThreePucksShouldMoveToTheSouthAndStoppedByTheWall(t *testing.T) {
	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#| |b| |#| | ",
		"#| | |#|b|#| | |#",
		"#|#| | |b| | |#|#",
		"#| | |#| |#| | |#",
		"#| |#| | | |#| |#",
		"#| | | |#| | | |#",
		"#################"}
	currentGame := game.Deserialize(rawGame)
	board := game.NewBoard()

	//WHEN
	currentGame = game.Tilt(currentGame, &board, "south")

	//THEN
	bluePuck33, bluePuck33Exists := currentGame.Pucks["3:3"]
	if !bluePuck33Exists || bluePuck33.Color != game.BLUE {
		t.Errorf("Puck should be on (3,3) %v", currentGame)
	}
	bluePuck34, bluePuck34Exists := currentGame.Pucks["3:4"]
	if !bluePuck34Exists || bluePuck34.Color != game.BLUE {
		t.Errorf("Puck should be on (3,4) %v", currentGame)
	}
	bluePuck35, bluePuck35Exists := currentGame.Pucks["3:5"]
	if !bluePuck35Exists || bluePuck35.Color != game.BLUE {
		t.Errorf("Puck should be on (3,5) %v", currentGame)
	}
}

func TestStuckPucksShouldNotMove(t *testing.T) {
	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|b|b|b|#| | ",
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
	bluePuck21, bluePuck21Exists := currentGame.Pucks["2:1"]
	if !bluePuck21Exists || bluePuck21.Color != game.BLUE {
		t.Errorf("Puck should be on (2,1) %v", currentGame)
	}
	bluePuck31, bluePuck31Exists := currentGame.Pucks["3:1"]
	if !bluePuck31Exists || bluePuck31.Color != game.BLUE {
		t.Errorf("Puck should be on (3,1) %v", currentGame)
	}
	bluePuck41, bluePuck41Exists := currentGame.Pucks["4:1"]
	if !bluePuck41Exists || bluePuck41.Color != game.BLUE {
		t.Errorf("Puck should be on (4,1) %v", currentGame)
	}
}

func TestSomePuckShouldMoveAndSomeShouldBeBlocked(t *testing.T) {
	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|b|b|b|#| | ",
		"#| | |#| |#| | |#",
		"#|#| | | | | |#|#",
		"#| | |#| |#| | |#",
		"#| |#| | | |#| |#",
		"#| | | |#| | | |#",
		"#################"}
	currentGame := game.Deserialize(rawGame)
	board := game.NewBoard()

	//WHEN
	currentGame = game.Tilt(currentGame, &board, "north")

	//THEN
	bluePuck20, bluePuck20Exists := currentGame.Pucks["2:0"]
	if !bluePuck20Exists || bluePuck20.Color != game.BLUE {
		t.Errorf("Puck should be on (2,0) %v", currentGame)
	}
	bluePuck31, bluePuck31Exists := currentGame.Pucks["3:1"]
	if !bluePuck31Exists || bluePuck31.Color != game.BLUE {
		t.Errorf("Puck should be on (3,1) %v", currentGame)
	}
	bluePuck40, bluePuck40Exists := currentGame.Pucks["4:0"]
	if !bluePuck40Exists || bluePuck40.Color != game.BLUE {
		t.Errorf("Puck should be on (4,0) %v", currentGame)
	}
}
