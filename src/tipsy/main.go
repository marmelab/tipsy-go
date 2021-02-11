package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tipsy/ai"
	"tipsy/game"
)

const (
	minMaxDepth = 3
)
func getNextMove(rw http.ResponseWriter, req *http.Request) {
    req.ParseForm()
    var currentGame game.Game
	err := json.NewDecoder(req.Body).Decode(&currentGame)
	if err != nil {
        http.Error(rw, err.Error(), http.StatusBadRequest)
        return
    }
	bestMove, _ := ai.GetNextMovesScores(currentGame, minMaxDepth, false)
    fmt.Fprintf(rw, "currentPlayer:"+bestMove)
}
func main() {
	http.HandleFunc("/", getNextMove)
    log.Fatal(http.ListenAndServe(":8082", nil))
}
