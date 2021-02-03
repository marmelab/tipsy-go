package tests

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"tipsy/ai"
	"tipsy/game"
)

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
	currentGame := game.Deserialize(rawGame)
	//WHEN
	state := ai.GetWinner(currentGame)
	//THEN
	if state != game.BLUE {
		t.Errorf("Winner should be blue: %v => %v", currentGame, state)
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
	rawGame := []string{game.BLUE,
		"#################",
		"#| | | |#| | | |#",
		"#| |#|b| |b|#| | ",
		"#| |r|#|r|#|r| |#",
		"#|#| |b| |b| |#|#",
		"#| |r|#|r|#|r| |#",
		"#| |#|b| |b|#| |#",
		"#| | | |#| | | |#",
		"#################"}
	currentGame := game.Deserialize(rawGame)
	//WHEN
	state := ai.GetWinner(currentGame)
	//THEN
	if state != game.BLUE {
		t.Errorf("Blue should win as he pushed the black puck out: %v => %v", currentGame, state)
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
