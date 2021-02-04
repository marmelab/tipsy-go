package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
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

	start := time.Now()
	bestMove, movesScore := ai.GetNextMovesScores(currentGame, minMaxDepth, *verbose)
	elapsed := time.Since(start)
	for move, moveScore := range movesScore {
		fmt.Printf("- %v => %v\n", move, moveScore)
	}
	fmt.Println()
	fmt.Printf("Found in %v\n\n", elapsed)
	fmt.Printf("Best Move : %v", bestMove)
}
