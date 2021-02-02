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
	for _, node := range (*board).Nodes {
		var rightPosition = [2]int{node.Position[0] + 1, node.Position[1]}
		var leftPosition = [2]int{node.Position[0] - 1, node.Position[1]}
		var upPosition = [2]int{node.Position[0], node.Position[1] - 1}
		var downPosition = [2]int{node.Position[0], node.Position[1] + 1}
		addEdge(&node, leftPosition, "left", board)
		addEdge(&node, rightPosition, "right", board)
		addEdge(&node, upPosition, "up", board)
		addEdge(&node, downPosition, "down", board)
	}
}

func addEdge(from *Node, to [2]int, value string, board *Board) {
	if Contains(to, (*board)) {
		var to = GetNode(to, board)
		(*board).Edges = append((*board).Edges, Edge{From: from, To: &to, Value: value})
	}
}

//GetNode gets the node given a position
func GetNode(position [2]int, board *Board) Node {
	for _, Node := range board.Nodes {
		if Node.Position[0] == position[0] && Node.Position[1] == position[1] {
			return Node
		}
	}
	return Node{}
}

// Contains return true if board contains a Vertex at a given position, and false otherwise.
func Contains(position [2]int, board Board) bool {
	for _, vertex := range board.Vertices {
		if vertex.Position[0] == position[0] && vertex.Position[1] == position[1] {
			return true
		}
	}
	return false
}
