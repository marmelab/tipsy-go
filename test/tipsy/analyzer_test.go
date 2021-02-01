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
	filePath := "./dataset/red-win.json"
	game := loadGame(filePath)
	//WHEN
	state := ai.GameState(game)
	//THEN
	if state != "red" {
		t.Errorf("Winner shoul be red: %v => %v", filePath, state)
	}
}

func TestBlueShouldBeTheWinnerIfAllSixBluePucksAreFlipped(t *testing.T) {

	//GIVEN
	filePath := "./dataset/blue-win.json"
	game := loadGame(filePath)
	//WHEN
	state := ai.GameState(game)
	//THEN
	if state != "blue" {
		t.Errorf("Winner should be blue: %v => %v", filePath, state)
	}
}

func TestShouldBeActiveIfNeitherBlueOrRedHaveSixPucksFlipped(t *testing.T) {

	//GIVEN
	filePath := "./dataset/active.json"
	game := loadGame(filePath)
	//WHEN
	state := ai.GameState(game)
	//THEN
	if state != "active" {
		t.Errorf("Game should be active: %v => %v", filePath, state)
	}
}

func TestBlueShouldWinWhenHeJustPushTheBlackPuckOut(t *testing.T) {

	//GIVEN
	filePath := "./dataset/blue-win-blackpuck.json"
	game := loadGame(filePath)
	//WHEN
	state := ai.GameState(game)
	//THEN
	if state != "blue" {
		t.Errorf("Blue should win as he pushed the black puck out: %v => %v", filePath, state)
	}
}

func TestRedShouldWinWhenHeJustPushTheBlackPuckOut(t *testing.T) {

	//GIVEN
	filePath := "./dataset/red-win-blackpuck.json"
	game := loadGame(filePath)

	//WHEN
	state := ai.GameState(game)

	//THEN
	if state != "red" {
		t.Errorf("Red should win as he pushed the black puck out: %v => %v", filePath, state)
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
