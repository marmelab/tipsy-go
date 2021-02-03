package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"tipsy/ai"
	"tipsy/game"
)

func main() {
	inputFilePath := flag.String("file",
		"./test/tipsy/dataset/active.json",
		"File with the board state, default is the starting board")
	askingPlayer := flag.String("player", game.BLUE, "Player asking for advices, default to 'blue'")
	verbose := flag.Bool("v", false, "Verbose output")
	file, err := os.Open(*inputFilePath)
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(file)

	var rawGame []string

	json.Unmarshal(byteValue, &rawGame)

	game := game.Deserialize(rawGame)

	moves := ai.GetNextMovesScores(game, *askingPlayer, *verbose)

	fmt.Printf("%v", moves)

}
