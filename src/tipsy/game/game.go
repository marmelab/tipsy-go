package game

import (
	"fmt"
	"strings"
	"tipsy/tools"
)

const (
	//BLUE color constant
	BLUE = "blue"
	//RED color constant
	RED = "red"
	//BLACK color constant
	BLACK = "black"
	//ACTIVE game is still active, no winner yet
	ACTIVE = "active"
	//RIGHT direction
	RIGHT = "right"
	//LEFT direction
	LEFT = "left"
	//UP direction
	UP = "up"
	//DOWN direction
	DOWN = "down"
)

//Directions of the game
var Directions = [4]string{RIGHT, LEFT, UP, DOWN}

//Game the game state
type Game struct {
	Pucks         map[string]Puck `json:"pucks"`
	FallenPucks   []Puck
	CurrentPlayer string `json:"currentPlayer"`
}

//SwitchPlayer return the opposite of currentPlayer
func SwitchPlayer(game Game) Game {
	if game.CurrentPlayer == RED {
		game.CurrentPlayer = BLUE
	} else {
		game.CurrentPlayer = RED
	}
	return game
}

//Deserialize a game represented in string array
func Deserialize(gameString []string) Game {
	var deserializedGame Game
	deserializedGame.Pucks = make(map[string]Puck)
	deserializedGame.CurrentPlayer = gameString[0]
	fmt.Println()
	fmt.Println(gameString[1])
	for rowIndex, line := range gameString[2 : BoardSize+2] {
		fmt.Println(strings.Replace(line, "|", " ", -1))
		characters := strings.Split(line, "|")

		for colIndex, char := range characters[1:] {
			position := [2]int{colIndex, rowIndex}
			switch char {
			case "r":
				deserializedGame.Pucks[tools.GetKeyFromPosition(position)] = Puck{Color: RED}
			case "R":
				deserializedGame.Pucks[tools.GetKeyFromPosition(position)] = Puck{Color: RED, Flipped: true}
			case "b":
				deserializedGame.Pucks[tools.GetKeyFromPosition(position)] = Puck{Color: BLUE}
			case "B":
				deserializedGame.Pucks[tools.GetKeyFromPosition(position)] = Puck{Color: BLUE, Flipped: true}
			case "x":
				deserializedGame.Pucks[tools.GetKeyFromPosition(position)] = Puck{Color: BLACK}
			}
		}
	}
	if len(gameString) == BoardSize+4 {
		for _, char := range strings.Split(gameString[10], "|") {
			switch char {
			case "r":
				deserializedGame.FallenPucks = append(deserializedGame.FallenPucks, Puck{Color: RED})
			case "b":
				deserializedGame.FallenPucks = append(deserializedGame.FallenPucks, Puck{Color: BLUE})
			case "x":
				deserializedGame.FallenPucks = append(deserializedGame.FallenPucks, Puck{Color: BLACK})
			}
		}
	}
	fmt.Println(gameString[BoardSize+2])
	fmt.Println()
	fmt.Println()
	return deserializedGame
}
