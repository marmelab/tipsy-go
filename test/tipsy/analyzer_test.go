package tests

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"tipsy/ai"
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
	game := game.Deserialize(rawGame)
	//THEN
	if game.BlackPuck.Position[0] != 3 || game.BlackPuck.Position[1] != 3 {
		t.Errorf("BlackPuck should be on (3,3) %v", game)
	}
	if game.RedPucks[0].Position[0] != 3 || game.RedPucks[0].Position[1] != 2 {
		t.Errorf("RedPuck should be on (3,2) %v", game)
	}
	if game.RedPucks[0].Flipped != true {
		t.Errorf("RedPuck should be flipped %v", game)
	}
	if game.BluePucks[0].Flipped == true {
		t.Errorf("BluePuck should not be flipped %v", game)
	}
	if game.BluePucks[0].Position[0] != 0 || game.BluePucks[0].Position[1] != 0 {
		t.Errorf("BluePuck should be on (0,0) %v", game)
	}
}
func TestRedShouldBeTheWinnerIfAllSixRedPucksAreFlipped(t *testing.T) {
	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|b| |b|#| | ",
		"#| |R|#|R|#|R| |#",
		"#|#| |b|x|b| |#|#",
		"#| |R|#|R|#|R| |#",
		"#| |#|b| |b|#| |#",
		"#| | | |#| | | |#",
		"#################"}
	game := game.Deserialize(rawGame)
	//WHEN
	state := ai.GetWinner(game)
	//THEN
	if state != "red" {
		t.Errorf("Winner shoul be red: %v => %v", game, state)
	}
}

func TestBlueShouldBeTheWinnerIfAllSixBluePucksAreFlipped(t *testing.T) {

	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|B| |B|#| | ",
		"#| |r|#|r|#|r| |#",
		"#|#| |B|x|B| |#|#",
		"#| |r|#|r|#|r| |#",
		"#| |#|B| |B|#| |#",
		"#| | | |#| | | |#",
		"#################"}
	game := game.Deserialize(rawGame)
	//WHEN
	state := ai.GetWinner(game)
	//THEN
	if state != "blue" {
		t.Errorf("Winner should be blue: %v => %v", game, state)
	}
}

func TestShouldBeActiveIfNeitherBlueOrRedHaveSixPucksFlipped(t *testing.T) {

	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|b| |b|#| | ",
		"#| |r|#|r|#|r| |#",
		"#|#| |b|x|b| |#|#",
		"#| |r|#|r|#|r| |#",
		"#| |#|b| |b|#| |#",
		"#| | | |#| | | |#",
		"#################"}
	game := game.Deserialize(rawGame)
	//WHEN
	state := ai.GetWinner(game)
	//THEN
	if state != "active" {
		t.Errorf("Game should be active: %v => %v", rawGame, state)
	}
}

func TestBlueShouldWinWhenHeJustPushTheBlackPuckOut(t *testing.T) {

	//GIVEN
	rawGame := []string{"blue",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|b| |b|#| | ",
		"#| |r|#|r|#|r| |#",
		"#|#| |b| |b| |#|#",
		"#| |r|#|r|#|r| |#",
		"#| |#|b| |b|#| |#",
		"#| | | |#| | | |#",
		"#################"}
	game := game.Deserialize(rawGame)
	//WHEN
	state := ai.GetWinner(game)
	//THEN
	if state != "blue" {
		t.Errorf("Blue should win as he pushed the black puck out: %v => %v", game, state)
	}
}

func TestRedShouldWinWhenHeJustPushTheBlackPuckOut(t *testing.T) {

	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|b| |b|#| | ",
		"#| |r|#|r|#|r| |#",
		"#|#| |b| |b| |#|#",
		"#| |r|#|r|#|r| |#",
		"#| |#|b| |b|#| |#",
		"#| | | |#| | | |#",
		"#################"}
	game := game.Deserialize(rawGame)

	//WHEN
	state := ai.GetWinner(game)

	//THEN
	if state != "red" {
		t.Errorf("Red should win as he pushed the black puck out: %v => %v", game, state)
	}
}

func loadGame(filePath string) game.Game {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(file)

	var game game.Game
	// we initialize our Users array

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &game)
	return game
}
