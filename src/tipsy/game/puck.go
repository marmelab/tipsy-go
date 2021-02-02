package game

//Puck representation, position and flipped or not
type Puck struct {
	Position [2]int `json:"position"`
	Flipped  bool   `json:"flipped"`
}
