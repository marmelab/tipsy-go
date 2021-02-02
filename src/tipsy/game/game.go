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
			switch char {
			case "r":
				game.Pucks = append(game.Pucks, Puck{Position: [2]int{col - 1, line - 2}, Color: "red"})
			case "R":
				game.Pucks = append(game.Pucks, Puck{Position: [2]int{col - 1, line - 2}, Color: "red", Flipped: true})
			case "b":
				game.Pucks = append(game.Pucks, Puck{Position: [2]int{col - 1, line - 2}, Color: "blue"})
			case "B":
				game.Pucks = append(game.Pucks, Puck{Position: [2]int{col - 1, line - 2}, Color: "blue", Flipped: true})
			case "x":
				game.Pucks = append(game.Pucks, Puck{Position: [2]int{col - 1, line - 2}, Color: "black"})
			}
		}
	}
	return game
}

func getNeighbor(position [2]int, board *Board, direction string) Node {
	puckNode := GetNode(position, board)
	return GetNodeTo(puckNode, board, direction)
}

func isAPuck(node Node, game Game) bool {
	for _, puck := range game.Pucks {
		if puck.Position == node.Position {
			return true
		}
	}
	return false
}
func getPuck(node Node, game Game) Puck {
	for _, puck := range game.Pucks {
		if puck.Position == node.Position {
			return puck
		}
	}
	panic("No Puck on this node")
}
func isAWall(node Node, board *Board) bool {
	return (Node{}) == GetNode(node.Position, board)
}

func getNextFreeCell(position [2]int, game Game, board *Board, direction string) [2]int {
	neighbor := getNeighbor(position, board, direction)
	if isAPuck(neighbor, game) || isAWall(neighbor, board) {
		return position
	}
	return getNextFreeCell(neighbor.Position, game, board, direction)
}

func movePuckTo(puck Puck, game Game, board *Board, direction string) []Puck {
	neighbor := getNeighbor(puck.Position, board, direction)
	var pucks []Puck
	if isAPuck(neighbor, game) {
		pucks = append(pucks, movePuckTo(getPuck(neighbor, game), game, board, direction)...)
	}
	nextFreeCell := getNextFreeCell(puck.Position, game, board, direction)
	puck.Position = nextFreeCell
	pucks = append(pucks, puck)
	return pucks

}

//Tilt the game in a given direction
func Tilt(game Game, board *Board, direction string) Game {
	var pucks []Puck
	for _, puck := range game.Pucks {
		pucks = append(pucks, movePuckTo(puck, game, board, direction)...)
	}
	game.Pucks = pucks
	return game
}
