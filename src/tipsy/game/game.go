package game

const (
	BLUE = "blue"
	RED  = "red"
)

type Game struct {
	ID          int
	Pucks       []Puck
	Players     [2]Player
	Board       Board
}

func New() Game {
	var game Game
	return game
}
