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
				game.Pucks[strconv.Itoa(col-1)+":"+strconv.Itoa(line-2)] = Puck{Color: "blue"}
			case "B":
				game.Pucks[strconv.Itoa(col-1)+":"+strconv.Itoa(line-2)] = Puck{Color: "blue", Flipped: true}
			case "x":
				game.Pucks[strconv.Itoa(col-1)+":"+strconv.Itoa(line-2)] = Puck{Color: "black"}
			}
		}
	}
	return game
}
func getPositionFromKey(key string) [2]int {
	positions := strings.Split(key, ":")
	x, _ := strconv.Atoi(positions[0])
	y, _ := strconv.Atoi(positions[1])
	return [2]int{x, y}
}
func getKeyFromPosition(position [2]int) string {
	return strconv.Itoa(position[0]) + ":" + strconv.Itoa(position[1])
}

func getNeighbor(position [2]int, board *Board, direction string) Node {
	puckNode := GetNode(position, board)
	return GetNodeTo(puckNode, board, direction)
}

func isAPuck(node Node, gamePucks map[string]Puck) bool {
	for key := range gamePucks {
		if getPositionFromKey(key) == node.Position {
			return true
		}
	}
	return false
}
func getPuck(node Node, gamePucks map[string]Puck) Puck {
	for key, puck := range gamePucks {
		if getPositionFromKey(key) == node.Position {
			return puck
		}
	}
	panic("No Puck on this node")
}
func isAWall(node Node, board *Board) bool {
	return (Node{}) == GetNode(node.Position, board)
}

func getNextFreeCell(position [2]int, gamePucks map[string]Puck, board *Board, direction string) [2]int {
	neighbor := getNeighbor(position, board, direction)
	if isAPuck(neighbor, gamePucks) || isAWall(neighbor, board) {
		return position
	}
	return getNextFreeCell(neighbor.Position, gamePucks, board, direction)
}

func movePuckTo(puckKey string, puck Puck, gamePucks map[string]Puck, board *Board, direction string) map[string]Puck {

	neighbor := getNeighbor(getPositionFromKey(puckKey), board, direction)
	var nodesWithPuck []Node
	for isAPuck(neighbor, gamePucks) {
		nodesWithPuck = append(nodesWithPuck, neighbor)
		neighbor = getNeighbor(neighbor.Position, board, direction)
	}

	pucks := make(map[string]Puck)
	for i := len(nodesWithPuck) - 1; i >= 0; i-- {
		nodeWithPuck := nodesWithPuck[i]
		nextFreeCell := getNextFreeCell(nodeWithPuck.Position, gamePucks, board, direction)
		puck := getPuck(nodeWithPuck, gamePucks)
		nextFreeCellKey := getKeyFromPosition(nextFreeCell)
		if nextFreeCell != nodeWithPuck.Position {
			pucks[nextFreeCellKey] = puck
			gamePucks[nextFreeCellKey] = puck
			delete(gamePucks, getKeyFromPosition(nodeWithPuck.Position))
		}
	}
	nextFreeCell := getNextFreeCell(getPositionFromKey(puckKey), gamePucks, board, direction)
	pucks[getKeyFromPosition(nextFreeCell)] = puck
	return pucks

}

//Tilt the game in a given direction
func Tilt(game Game, board *Board, direction string) Game {
	pucks := make(map[string]Puck)
	for key, puck := range game.Pucks {
		movedPucks := movePuckTo(key, puck, game.Pucks, board, direction)
		for key, puck := range movedPucks {
			pucks[key] = puck
		}
	}
	game.Pucks = pucks
	return game
}
