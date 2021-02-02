package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"tipsy/ai"
	"tipsy/game"
)

func main() {
	inputFilePath := os.Args[1]
	file, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(file)

	var game game.Game

	json.Unmarshal(byteValue, &game)

	fmt.Println(ai.GetWinner(game))
}
