package tests

import (
	"testing"
	"tipsy/game"
)

func TestGameShouldBeWellDesiaralized(t *testing.T) {
	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#|b| | |#| | | |#",
		"#| |#| | | |#| | ",
		"#| | |#|R|#| | |#",
		"#|#| | |x| | |#|#",
		"#| | |#| |#| | |#",
		"#| |#| | | |#| |#",
		"#| | | |#| | | |#",
		"#################"}
	//WHEN
	game := game.Deserialize(rawGame)
	//THEN
	if game.Pucks[2].Position[0] != 3 || game.Pucks[2].Position[1] != 3 {
		t.Errorf("BlackPuck should be on (3,3) %v", game)
	}
	if game.Pucks[1].Position[0] != 3 || game.Pucks[1].Position[1] != 2 || game.Pucks[1].Color != "red" {
		t.Errorf("RedPuck should be on (3,2) %v", game)
	}
	if game.Pucks[1].Flipped != true {
		t.Errorf("RedPuck should be flipped %v", game)
	}
	if game.Pucks[0].Position[0] != 0 || game.Pucks[0].Position[1] != 0 || game.Pucks[0].Color != "blue" {
		t.Errorf("BluePuck should be on (0,0) %v", game)
	}
	if game.Pucks[0].Flipped != false {
		t.Errorf("BluePuck should not be flipped %v", game)
	}
}
