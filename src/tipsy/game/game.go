package game

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
	RedPucks      [6]Puck `json:"redPucks"`
	BluePucks     [6]Puck `json:"bluePucks"`
	BlackPuck     Puck    `json:"blackPuck"`
	CurrentPlayer string  `json:"currentPlayer"`
}
