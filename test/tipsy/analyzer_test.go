package tests

import (
	"testing"
	"os"
	"io/ioutil"
	"tipsy/game"
	"tipsy/ai"
	"encoding/json"
)

func TestGameInit(t *testing.T) {

	var game game.Game

	loadGame("./dataset/win.json", game)
	if ai.GameState(game) != "red"{
		t.Errorf("Winner shoul be red")
	}
}

func loadGame(filePath string, game game.Game){
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(file)

	// we initialize our Users array

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &game)
}
