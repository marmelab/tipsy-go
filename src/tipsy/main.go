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

	var rawGame []string

	json.Unmarshal(byteValue, &rawGame)

	game := game.Deserialize(rawGame)

	moves := ai.GetNextMoves(game)

	fmt.Printf("%v", moves)

}
