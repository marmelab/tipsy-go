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

const (
	minMaxDepth = 3
)

func main() {
	inputFilePath := flag.String("file",
		"./test/tipsy/dataset/active.json",
		"File with the board state, default is the starting board")
	verbose := flag.Bool("v", false, "Verbose output")
	flag.Parse()
	file, err := os.Open(*inputFilePath)
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(file)

	var rawGame []string

	json.Unmarshal(byteValue, &rawGame)

	currentGame := game.Deserialize(rawGame)
	bestMove := ai.GetNextMovesScores(currentGame, minMaxDepth, *verbose)
	fmt.Printf("Best Move : %v", bestMove)
}
