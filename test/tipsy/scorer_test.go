package tests

import (
	"testing"
	"tipsy/ai"
	"tipsy/game"
)

func TestRedShouldBeTheWinnerIfAllSixRedPucksAreFlipped(t *testing.T) {
	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|b| |b|#| | ",
		"#| |R|#|R|#|R| |#",
		"#|#| |b|x|b| |#|#",
		"#| |R|#|R|#|R| |#",
		"#| |#|b| |b|#| |#",
		"#| | | |#| | | |#",
		"#################"}
	game := game.Deserialize(rawGame)
	//WHEN
	score := ai.GetScore(game)
	//THEN
	if score != ai.WinningScore {
		t.Errorf("Winner shoul be red: %v => %v", game, score)
	}
}

func TestBlueShouldLoseIfAllSixBluePucksAreFlipped(t *testing.T) {

	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|B| |B|#| | ",
		"#| |r|#|r|#|r| |#",
		"#|#| |B|x|B| |#|#",
		"#| |r|#|r|#|r| |#",
		"#| |#|B| |B|#| |#",
		"#| | | |#| | | |#",
		"#################"}
	currentGame := game.Deserialize(rawGame)
	//WHEN
	score := ai.GetScore(currentGame)
	//THEN
	if score != ai.LosingScore {
		t.Errorf("Winner should be blue: %v", score)
	}
}

func TestShouldBeActiveIfNeitherBlueOrRedHaveSixPucksFlipped(t *testing.T) {

	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|b| |b|#| | ",
		"#| |r|#|r|#|r| |#",
		"#|#| |b|x|b| |#|#",
		"#| |r|#|r|#|r| |#",
		"#| |#|b| |b|#| |#",
		"#| | | |#| | | |#",
		"#################"}
	game := game.Deserialize(rawGame)
	//WHEN
	score := ai.GetScore(game)
	//THEN
	if score == ai.WinningScore || score == ai.LosingScore {
		t.Errorf("Game should be active: %v ", score)
	}
}

func TestBlueShouldWinWhenHeJustPushTheBlackPuckOut(t *testing.T) {

	//GIVEN
	rawGame := []string{game.BLUE,
		"#################",
		"#| | | |#| | | |#",
		"#| |#|b| |b|#| | ",
		"#| |r|#|r|#|r| |#",
		"#|#| |b| |b| |#|#",
		"#| |r|#|r|#|r| |#",
		"#| |#|b| |b|#| |#",
		"#| | | |#| | | |#",
		"#################",
		"x"}
	currentGame := game.Deserialize(rawGame)
	//WHEN
	score := ai.GetScore(currentGame)
	//THEN
	if score != ai.WinningScore {
		t.Errorf("Blue should win as he pushed the black puck out: %v => %v", currentGame, score)
	}
}

func TestRedShouldWinWhenHeJustPushTheBlackPuckOut(t *testing.T) {

	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|b| |b|#| | ",
		"#| |r|#|r|#|r| |#",
		"#|#| |b| |b| |#|#",
		"#| |r|#|r|#|r| |#",
		"#| |#|b| |b|#| |#",
		"#| | | |#| | | |#",
		"#################",
		"x"}
	game := game.Deserialize(rawGame)

	//WHEN
	score := ai.GetScore(game)

	//THEN
	if score != ai.WinningScore {
		t.Errorf("Red should win as he pushed the black puck out: %v => %v", game, score)
	}
}

func TestRedShouldWinWhenHeJustPushHisLastUnFlippedPuckOut(t *testing.T) {

	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|b| |b|#| | ",
		"#| |R|#|R|#| | |#",
		"#|#| |b| |b| |#|#",
		"#| |R|#|R|#|R| |#",
		"#| |#|b| |b|#| |#",
		"#| | | |#| | | |#",
		"#################",
		"r"}
	game := game.Deserialize(rawGame)

	//WHEN
	score := ai.GetScore(game)

	//THEN
	if score != ai.WinningScore {
		t.Errorf("Red should win as he pushed his last unflipped puck out: %v => %v", game, score)
	}
}
func TestRedShouldWinWhenHeJustPushHisTwoLastUnFlippedPucksOut(t *testing.T) {

	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|b| |b|#| | ",
		"#| |R|#| |#| | |#",
		"#|#| |b| |b| |#|#",
		"#| |R|#|R|#|R| |#",
		"#| |#|b| |b|#| |#",
		"#| | | |#| | | |#",
		"#################",
		"r|r"}
	game := game.Deserialize(rawGame)

	//WHEN
	score := ai.GetScore(game)

	//THEN
	if score != ai.WinningScore {
		t.Errorf("Red should win as he pushed his two last unflipped pucks out: %v => %v", game, score)
	}
}
func TestRedShouldLoseWhenRedJustPushLastBlueUnFlippedPucksOut(t *testing.T) {

	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|B| |B|#| | ",
		"#| |r|#|r|#|r| |#",
		"#|#| |B| |B| |#|#",
		"#| |r|#|r|#|r| |#",
		"#| |#|B| | |#| |#",
		"#| | | |#| | | |#",
		"#################",
		"b"}
	game := game.Deserialize(rawGame)

	//WHEN
	score := ai.GetScore(game)

	//THEN
	if score != ai.LosingScore {
		t.Errorf("Red should lose as Red pushed Blue last unflipped puck out: %v => %v", game, score)
	}
}

func TestRedShouldLoseWhenRedJustPushLastBlueAndLastRedUnFlippedPucksOut(t *testing.T) {

	//GIVEN
	rawGame := []string{"red",
		"#################",
		"#| | | |#| | | |#",
		"#| |#|B| |B|#| | ",
		"#| |R|#|R|#|R| |#",
		"#|#| |B| |B| |#|#",
		"#| |R|#|R|#| | |#",
		"#| |#|B| | |#| |#",
		"#| | | |#| | | |#",
		"#################",
		"b|r"}
	game := game.Deserialize(rawGame)

	//WHEN
	score := ai.GetScore(game)

	//THEN
	if score != ai.LosingScore {
		t.Errorf("Blue should win as Red pushed Blue and Red last unflipped pucks out: %v => %v", game, score)
	}
}
