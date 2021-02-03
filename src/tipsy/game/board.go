package game

import (
	"tipsy/tools"
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
		addEdge(node, leftPosition, "west", board)
		addEdge(node, rightPosition, "east", board)
		addEdge(node, upPosition, "north", board)
		addEdge(node, downPosition, "south", board)
	}
}

func addEdge(from Node, to [2]int, value string, board *Board) {
	if Contains(to, board) {
		var to = GetNode(to, board)
		(*board).Edges = append(board.Edges, Edge{From: from, To: to, Value: value})
	}
}

//GetNode gets the node given a position
func GetNode(position [2]int, board *Board) Node {
	for _, node := range board.Nodes {
		if node.Position[0] == position[0] && node.Position[1] == position[1] {
			return node
		}
	}
	return Node{}
}

//GetNodeTo gets the node from a given one to a given direction
func GetNodeTo(node Node, board *Board, direction string) Node {
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
