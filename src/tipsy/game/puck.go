package game

type Puck struct {
	Position [2]int `json:"position"`
	Color    string `json:"color"`
	Flipped  bool   `json:"flipped"`
}
