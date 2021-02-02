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
	RedPucks      []Puck `json:"redPucks"`
	BluePucks     []Puck `json:"bluePucks"`
	BlackPuck     Puck   `json:"blackPuck"`
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
				game.RedPucks = append(game.RedPucks, Puck{Position: [2]int{col - 1, line - 2}})
			}
			if char == "R" {
				game.RedPucks = append(game.RedPucks, Puck{Position: [2]int{col - 1, line - 2}, Flipped: true})
			}
			if char == "b" {
				game.BluePucks = append(game.BluePucks, Puck{Position: [2]int{col - 1, line - 2}})
			}
			if char == "B" {
				game.BluePucks = append(game.BluePucks, Puck{Position: [2]int{col - 1, line - 2}, Flipped: true})
			}
			if char == "x" {
				game.BlackPuck = Puck{Position: [2]int{col - 1, line - 2}}
			}
		}
	}
	return game
}
