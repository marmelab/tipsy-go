package game

type Puck struct {
	Position [2]int `json:"position"`
	Flipped  bool   `json:"flipped"`
}
