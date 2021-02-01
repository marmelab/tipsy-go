package game

const (
	BLUE = "blue"
	RED  = "red"
)

type Game struct {
	RedPucks  []Puck `json:"redPucks"`
	BluePucks []Puck `json:"bluePucks"`
	BlackPuck Puck   `json:"blackPuck"`
	CurrentPlayer string `json:"currentPlayer"`
}

func New() Game {
	var game Game
	return game
}
