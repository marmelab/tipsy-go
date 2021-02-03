package game

import (
	"fmt"
	"strconv"
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
	Pucks         map[string]Puck `json:"pucks"`
	CurrentPlayer string          `json:"currentPlayer"`
}

//Deserialize a game represented in string array
func Deserialize(gameString []string) Game {
	var game Game
	game.Pucks = make(map[string]Puck)
	for line, value := range gameString {
		if line == 0 {
			game.CurrentPlayer = value
		}
		characters := strings.Split(gameString[line], "|")
		fmt.Println(characters)

		for col, char := range characters {
			switch char {
			case "r":
				game.Pucks[strconv.Itoa(col-1)+":"+strconv.Itoa(line-2)] = Puck{Color: "red"}
			case "R":
				game.Pucks[strconv.Itoa(col-1)+":"+strconv.Itoa(line-2)] = Puck{Color: "red", Flipped: true}
			case "b":
				game.Pucks[strconv.Itoa(col-1)+":"+strconv.Itoa(line-2)] = Puck{Color: BLUE}
			case "B":
				game.Pucks[strconv.Itoa(col-1)+":"+strconv.Itoa(line-2)] = Puck{Color: BLUE, Flipped: true}
			case "x":
				game.Pucks[strconv.Itoa(col-1)+":"+strconv.Itoa(line-2)] = Puck{Color: "black"}
			}
		}
	}
	return game
}

