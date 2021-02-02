package game

//Puck representation, position and flipped or not
type Puck struct {
	Position [2]int `json:"position"`
	Color    string `json:"color"`
	Flipped  bool   `json:"flipped"`
}
