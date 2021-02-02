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

func getNeighbor(puck Puck, board *Board, direction string) *Node {
	puckNode := GetNode(puck.Position, board)
	return GetNodeTo(&puckNode, board, direction)
}

func isAPuck(node *Node, game *Game) bool {
	for _, puck := range game.Pucks {
		fmt.Println(puck)
	}
	return true
}
func getPuck(node *Node, game *Game) *Puck {
	for _, puck := range game.Pucks {
		if puck.Position[0] == node.Position[0] && puck.Position[1] == node.Position[1] {
			return &puck
		}
	}
	panic("No Puck on this node")
}

func getNextFreeCell(puck *Puck, game *Game, board *Board, direction string) [2]int {
	neighbor := getNeighbor(*puck, board, direction)
	if isAPuck(neighbor, game) {
		return puck.Position
	}
	return getNextFreeCell(getPuck(neighbor, game), game, board, direction)
}

func movePuckTo(puck *Puck, game *Game, board *Board, direction string) {
	neighbor := getNeighbor(*puck, board, direction)
	//if a puck, move it
	if isAPuck(neighbor, game) {
		movePuckTo(getPuck(neighbor, game), game, board, direction)
	} else {
		nextFreeCel := getNextFreeCell(puck, game, board, direction)
		puck.Position = nextFreeCel
	}
}

//Tilt the game in a given direction
func Tilt(game *Game, board *Board, direction string) {
	for _, puck := range game.Pucks {
		movePuckTo(&puck, game, board, direction)
	}
}
