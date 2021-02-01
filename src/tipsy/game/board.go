package game

import (
	"tipsy/tools"
)

//Board : the board of tipsy game
type Board struct {
	Vertices []Vertex
	Edges    []Edge
}

//NewBoard initialize an empty board with obstacles and exits
func NewBoard() Board {

	var board Board
	initVertices(&board)
	initEdges(&board)
	return board
}

func initVertices(board *Board) {
	obstacles := [][2]int{
		{0, 3}, {1, 1}, {1, 5}, {2, 2},
		{2, 4}, {3, 0}, {3, 6}, {4, 2},
		{4, 4}, {5, 1}, {5, 5}, {6, 3}}
	exits := [][2]int{
		{1, -1}, {7, 1}, {-1, 5}, {5, 7}}

	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			if !tools.ArrayContains(obstacles, []int{i, j}) {
				(*board).Vertices = append((*board).Vertices, Vertex{Position: [2]int{i, j}})
			}
		}
	}
	for _, exit := range exits {
		(*board).Vertices = append((*board).Vertices, Vertex{Position: [2]int{exit[0], exit[1]}, Exit: true})
	}
}

func initEdges(board *Board) {
	for _, vertex := range (*board).Vertices {
		var rightPosition = [2]int{vertex.Position[0] + 1, vertex.Position[1]}
		var leftPosition = [2]int{vertex.Position[0] - 1, vertex.Position[1]}
		var upPosition = [2]int{vertex.Position[0], vertex.Position[1] - 1}
		var downPosition = [2]int{vertex.Position[0], vertex.Position[1] + 1}
		addEdge(&vertex, leftPosition, "left", board)
		addEdge(&vertex, rightPosition, "right", board)
		addEdge(&vertex, upPosition, "up", board)
		addEdge(&vertex, downPosition, "down", board)
	}
}

func addEdge(from *Vertex, to [2]int, value string, board *Board) {
	if Contains(to, (*board)) {
		var to = getVertex(to, (*board))
		(*board).Edges = append((*board).Edges, Edge{From: from, To: &to, Value: value})
	}
}
func getVertex(position [2]int, board Board) Vertex {
	for _, vertex := range board.Vertices {
		if vertex.Position[0] == position[0] && vertex.Position[1] == position[1] {
			return vertex
		}
	}
	return Vertex{}
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
