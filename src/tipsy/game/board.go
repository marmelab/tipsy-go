package game

import (
	"tipsy/tools"
)

type Board struct {
	Vertices []Vertex
	Edges    []Edge
}

func NewBoard() Board {

	var board Board
	initVertices(&board)
	return board
}

func initVertices(board *Board) {
	obstacles := [][]int{
		{0, 3}, {1, 1}, {1, 5}, {2, 2},
		{2, 4}, {3, 0}, {3, 6}, {4, 2},
		{4, 4}, {5, 1}, {5, 5}, {6, 3}}
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			if !tools.ArrayContains(obstacles, []int{i, j}) {
				(*board).Vertices = append((*board).Vertices, Vertex{Position: [2]int{i, j}})
			}
		}
	}
}

func Contains(position []int, board Board) bool {
	for _, vertex := range board.Vertices {
		if vertex.Position[0] == position[0] && vertex.Position[1] == position[1] {
			return true
		}
	}
	return false
}
