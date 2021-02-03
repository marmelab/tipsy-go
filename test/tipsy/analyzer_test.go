package tests

import (
	"encoding/json"
	"io/ioutil"
	"os"
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
	state := ai.GetScore(game, "blue")
	//THEN
	if state != ai.LosingScore {
		t.Errorf("Winner shoul be red: %v => %v", game, state)
	}
}

func TestBlueShouldBeTheWinnerIfAllSixBluePucksAreFlipped(t *testing.T) {

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
	state := ai.GetScore(currentGame, "blue")
	//THEN
	if state != ai.WinningScore {
		t.Errorf("Winner should be blue: %v => %v", currentGame, state)
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
	state := ai.GetScore(game, "blue")
	//THEN
	if state != ai.ActiveScore {
		t.Errorf("Game should be active: %v => %v", rawGame, state)
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
	state := ai.GetScore(currentGame, "blue")
	//THEN
	if state != ai.WinningScore {
		t.Errorf("Blue should win as he pushed the black puck out: %v => %v", currentGame, state)
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
	state := ai.GetScore(game, "red")

	//THEN
	if state != ai.WinningScore {
		t.Errorf("Red should win as he pushed the black puck out: %v => %v", game, state)
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
	state := ai.GetScore(game, "red")

	//THEN
	if state != ai.WinningScore {
		t.Errorf("Red should win as he pushed his last unflipped puck out: %v => %v", game, state)
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
	state := ai.GetScore(game, "red")

	//THEN
	if state != ai.WinningScore {
		t.Errorf("Red should win as he pushed his two last unflipped pucks out: %v => %v", game, state)
	}
}
func TestBlueShouldWinWhenRedJustPushLastBlueUnFlippedPucksOut(t *testing.T) {

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
	state := ai.GetScore(game, "blue")

	//THEN
	if state != ai.WinningScore {
		t.Errorf("Blue should win as Red pushed Blue last unflipped puck out: %v => %v", game, state)
	}
}

func TestBlueShouldWinWhenRedJustPushLastBlueAndLastRedUnFlippedPucksOut(t *testing.T) {

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
	state := ai.GetScore(game, "blue")

	//THEN
	if state != ai.WinningScore {
		t.Errorf("Blue should win as Red pushed Blue and Red last unflipped pucks out: %v => %v", game, state)
	}
}

//last blue out in 1 move
func TestMoveToEastShouldBePartOfTheWinsWhenLastBlueNearExit(t *testing.T) {

	//GIVEN
	rawGame := []string{"blue",
		"#### ############",
		"#| | | |#| | | |#",
		"#| |#| |B|B|#|b| ",
		"#| |R|#|R|#| |R|#",
		"#|#| | | |B|B|#|#",
		"#| |R|#|R|#| |B|#",
		" | |#| | | |#| |#",
		"#| | | |#| | | |#",
		"############ ####"}
	game := game.Deserialize(rawGame)

	//WHEN
	moves := ai.GetNextMovesScores(game, "blue")

	//THEN
	if len(moves) != 3 {
		t.Errorf("Analyzer should return just one winning move %v", moves)
	}
	rightWin := moves["right"]
	downRightWin := moves["down:right"]
	leftRightWin := moves["left:right"]
	if !(rightWin == ai.WinningScore) || !(downRightWin == ai.WinningScore) || !(leftRightWin == ai.WinningScore) {
		t.Errorf("The winning move should be 'right' %v", moves)
	}
}

//last blue out in 1 move when red play
func TestMoveToEastShouldBePartOfTheLoseWhenLastBlueNearExit(t *testing.T) {

	//GIVEN
	rawGame := []string{"red",
		"#### ############",
		"#| | | |#| | | |#",
		"#| |#| |B|B|#|b| ",
		"#| |R|#|R|#| |R|#",
		"#|#| | | |B|B|#|#",
		"#| |R|#|R|#| |B|#",
		" | |#| | | |#| |#",
		"#| | | |#| | | |#",
		"############ ####"}
	game := game.Deserialize(rawGame)

	//WHEN
	moves := ai.GetNextMovesScores(game, "red")

	//THEN
	if len(moves) != 3 {
		t.Errorf("Analyzer should return just one winning move %v", moves)
	}
	rightWin := moves["right"]
	downRightWin := moves["down:right"]
	leftRightWin := moves["left:right"]
	if (rightWin == ai.WinningScore) || (downRightWin == ai.WinningScore) || (leftRightWin == ai.WinningScore) {
		t.Errorf("The losing move should be 'right' %v", moves)
	}
}

//one blue out in 2 moves
func TestMoveToUpRightShouldBePartOfTheWinsWhenLastBlueIsOneCellNearExit(t *testing.T) {

	//GIVEN
	rawGame := []string{"blue",
		"#### ############",
		"#| | | |#| | | |#",
		"#| |#|B| |B|#|R| ",
		"#| | |#| |#| |b|#",
		"#|#|R|B|r|B| |#|#",
		"#| |R|#|x|#| | |#",
		" | |#| |r| |#| |#",
		"#| | | |#|B|r| |#",
		"############ ####"}
	game := game.Deserialize(rawGame)

	//WHEN
	moves := ai.GetNextMovesScores(game, "blue")

	//THEN
	upRightWin := moves["up:right"]
	if !(upRightWin == ai.WinningScore) {
		t.Errorf("The winning move should be 'right' %v", moves)
	}
}

//two blue out in 2 moves
//one blue one red out => should not win
//black out in one move
//black out in one two move
//black out and red out in two move
func loadGame(filePath string) game.Game {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(file)

	var game game.Game
	// we initialize our Users array

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &game)
	return game
}
