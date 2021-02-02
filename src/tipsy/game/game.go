package game

import (
	"fmt"
	"strings"
)

const (
	//BLUE color constant
	BLUE = "blue"
	//RED color constant
	RED = "red"
	//ACTIVE game is still active, no winner yet
	ACTIVE = "active"
)

//Game the game state
type Game struct {
	Pucks         []Puck `json:"pucks"`
	CurrentPlayer string `json:"currentPlayer"`
}

//Deserialize a game represented in string array
func Deserialize(gameString []string) Game {
	var game Game
	for line, value := range gameString {
		if line == 0 {
			game.CurrentPlayer = value
		}
		characters := strings.Split(gameString[line], "|")
		fmt.Println(characters)
		for col, char := range characters {
			if char == "r" {
				game.Pucks = append(game.Pucks, Puck{Position: [2]int{col - 1, line - 2}, Color: "red"})
			}
			if char == "R" {
				game.Pucks = append(game.Pucks, Puck{Position: [2]int{col - 1, line - 2}, Color: "red", Flipped: true})
			}
			if char == "b" {
				game.Pucks = append(game.Pucks, Puck{Position: [2]int{col - 1, line - 2}, Color: "blue"})
			}
			if char == "B" {
				game.Pucks = append(game.Pucks, Puck{Position: [2]int{col - 1, line - 2}, Color: "blue", Flipped: true})
			}
			if char == "x" {
				game.Pucks = append(game.Pucks, Puck{Position: [2]int{col - 1, line - 2}, Color: "black"})
			}
		}
	}
	return game
}
