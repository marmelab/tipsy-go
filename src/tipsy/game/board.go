package game

import (
	"tipsy/tools"
)

const (
	//BoardSize the size of the board
	BoardSize = 7
)

//Board : the board of tipsy game
type Board struct {
	Nodes []Node
	Edges []Edge
}

//NewBoard initialize an empty board with obstacles and exits
func NewBoard() Board {

	var board Board
	initNodes(&board)
	initEdges(&board)
	return board
}

func initNodes(board *Board) {
	obstacles := [][2]int{
		{0, 3}, {1, 1}, {1, 5}, {2, 2},
		{2, 4}, {3, 0}, {3, 6}, {4, 2},
		{4, 4}, {5, 1}, {5, 5}, {6, 3}}
	exits := [][2]int{
		{1, -1}, {7, 1}, {-1, 5}, {5, 7}}

	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			if !tools.ArrayContains(obstacles, []int{i, j}) {
				(*board).Nodes = append((*board).Nodes, Node{Position: [2]int{i, j}})
			}
		}
	}
	for _, exit := range exits {
		(*board).Nodes = append((*board).Nodes, Node{Position: [2]int{exit[0], exit[1]}, Exit: true})
	}
}

func initEdges(board *Board) {
	for _, node := range board.Nodes {
		var rightPosition = [2]int{node.Position[0] + 1, node.Position[1]}
		var leftPosition = [2]int{node.Position[0] - 1, node.Position[1]}
		var upPosition = [2]int{node.Position[0], node.Position[1] - 1}
		var downPosition = [2]int{node.Position[0], node.Position[1] + 1}
		addEdge(node, leftPosition, LEFT, board)
		addEdge(node, rightPosition, RIGHT, board)
		addEdge(node, upPosition, UP, board)
		addEdge(node, downPosition, DOWN, board)
	}
}

func addEdge(from Node, to [2]int, value string, board *Board) {
	if Contains(to, board) {
		var to = getNode(to, board)
		(*board).Edges = append(board.Edges, Edge{From: from, To: to, Value: value})
	}
}

func getNode(position [2]int, board *Board) Node {
	for _, node := range board.Nodes {
		if node.Position[0] == position[0] && node.Position[1] == position[1] {
			return node
		}
	}
	return Node{}
}

func getNodeTo(node Node, board *Board, direction string) Node {
	for _, edge := range board.Edges {
		if edge.From == node && edge.Value == direction {
			return edge.To
		}
	}
	return Node{}
}

// Contains return true if board contains a Node at a given position, and false otherwise.
func Contains(position [2]int, board *Board) bool {
	for _, Node := range board.Nodes {
		if Node.Position[0] == position[0] && Node.Position[1] == position[1] {
			return true
		}
	}
	return false
}

func getNeighbor(position [2]int, board *Board, direction string) Node {
	puckNode := getNode(position, board)
	return getNodeTo(puckNode, board, direction)
}

func isAPuck(node Node, gamePucks map[string]Puck) bool {
	for key := range gamePucks {
		if tools.GetPositionFromKey(key) == node.Position {
			return true
		}
	}
	return false
}
func getPuck(node Node, gamePucks map[string]Puck) Puck {
	for key, puck := range gamePucks {
		if tools.GetPositionFromKey(key) == node.Position {
			return puck
		}
	}
	panic("No Puck on this node")
}
func isAWall(node Node, board *Board) bool {
	return (Node{}) == getNode(node.Position, board)
}

func getNextFreeCell(position [2]int, gamePucks map[string]Puck, board *Board, direction string) [2]int {
	neighbor := getNeighbor(position, board, direction)
	if isAPuck(neighbor, gamePucks) || isAWall(neighbor, board) {
		return position
	}
	return getNextFreeCell(neighbor.Position, gamePucks, board, direction)
}

func isExit(position [2]int, board *Board) bool {
	node := getNode(position, board)
	return node.Exit
}

func movePuckTo(puckKey string, currentPuck Puck,
	gamePucks map[string]Puck, board *Board, direction string) (map[string]Puck, []Puck) {

	neighbor := getNeighbor(tools.GetPositionFromKey(puckKey), board, direction)
	var nodesWithPuck []Node
	for isAPuck(neighbor, gamePucks) {
		nodesWithPuck = append(nodesWithPuck, neighbor)
		neighbor = getNeighbor(neighbor.Position, board, direction)
	}

	pucks := make(map[string]Puck)
	var fallenPucks []Puck
	for i := len(nodesWithPuck) - 1; i >= 0; i-- {
		nodeWithPuck := nodesWithPuck[i]
		nextFreeCell := getNextFreeCell(nodeWithPuck.Position, gamePucks, board, direction)
		puck := getPuck(nodeWithPuck, gamePucks)
		nextFreeCellKey := tools.GetKeyFromPosition(nextFreeCell)
		if nextFreeCell != nodeWithPuck.Position {
			if isExit(nextFreeCell, board) {
				fallenPucks = append(fallenPucks, puck)
			} else {
				pucks[nextFreeCellKey] = puck
				gamePucks[nextFreeCellKey] = puck
			}
			delete(gamePucks, tools.GetKeyFromPosition(nodeWithPuck.Position))
		}
	}
	nextFreeCell := getNextFreeCell(tools.GetPositionFromKey(puckKey), gamePucks, board, direction)
	if isExit(nextFreeCell, board) {
		fallenPucks = append(fallenPucks, currentPuck)
	} else {
		pucks[tools.GetKeyFromPosition(nextFreeCell)] = currentPuck
	}
	return pucks, fallenPucks

}

//Tilt the game in a given direction
func Tilt(game Game, board *Board, direction string) Game {
	gamePucks := make(map[string]Puck)
	var gameFallenPucks []Puck
	for key, puck := range game.Pucks {
		movedPucks, fallenPucks := movePuckTo(key, puck, game.Pucks, board, direction)
		for key, puck := range movedPucks {
			gamePucks[key] = puck
		}
		gameFallenPucks = append(gameFallenPucks, fallenPucks...)
	}
	game.Pucks = gamePucks
	game.FallenPucks = gameFallenPucks
	return game
}
