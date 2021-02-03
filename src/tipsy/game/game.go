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
)

//Game the game state
type Game struct {
	Pucks         map[string]Puck `json:"pucks"`
	FallenPucks   []Puck
	CurrentPlayer string `json:"currentPlayer"`
}

//Deserialize a game represented in string array
func Deserialize(gameString []string) Game {
	var deserializedGame Game
	deserializedGame.Pucks = make(map[string]Puck)
	deserializedGame.CurrentPlayer = gameString[0]
	for rowIndex, line := range gameString[2 : BOARD_SIZE+1] {
		fmt.Println(line)
		characters := strings.Split(line, "|")
		fmt.Println(characters)

		for colIndex, char := range characters[1 : len(characters)-1] {
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
	if len(gameString) == BOARD_SIZE+4 {
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
	return deserializedGame
}
