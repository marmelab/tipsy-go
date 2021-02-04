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
	currentGame := game.Deserialize(rawGame)

	//WHEN
	moves := ai.GetNextMovesScores(currentGame, true)

	//THEN
	rightWin := moves[game.RIGHT]
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
	currentGame := game.Deserialize(rawGame)

	//WHEN
	moves := ai.GetNextMovesScores(currentGame, true)

	//THEN
	rightWin := moves[game.RIGHT]
	downRightWin := moves["down:right"]
	leftRightWin := moves["left:right"]
	if (rightWin != ai.LosingScore) || (downRightWin != ai.LosingScore) || (leftRightWin != ai.LosingScore) {
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
	moves := ai.GetNextMovesScores(game, true)

	//THEN
	upRightWin := moves["up:right"]
	if !(upRightWin == ai.WinningScore) {
		t.Errorf("The winning move should be 'right' %v", moves)
	}
}

func TestBestMoveShouldBeToMoveBlackFarOfTheExit(t *testing.T) {

	//GIVEN
	rawGame := []string{"blue",
		"#### ############",
		"#| | | |#| | |x|#",
		"#| |#| | | |#| | ",
		"#| | |#| |#| | |#",
		"#|#| | | | | |#|#",
		"#| | |#| |#| | |#",
		" | |#| | | |#| |#",
		"#| | | |#| | | |#",
		"############ ####"}
	game := game.Deserialize(rawGame)

	//WHEN
	moves := ai.GetNextMovesScores(game, true)

	//THEN
	leftDown := moves["left:down"]

	if leftDown != -400 {
		t.Errorf("Best move should be 'left:down' %v", moves)
	}
}

func TestMoveToDownRightShouldBeTheBestMove(t *testing.T) {

	//GIVEN
	rawGame := []string{"blue",
		"#### ############",
		"#| | | |#| | |b|#",
		"#| |#| |B|B|#| | ",
		"#| |R|#|R|#| |R|#",
		"#|#| | | |B|B|#|#",
		"#| |R|#|R|#|r|B|#",
		" | |#| | | |#| |#",
		"#| | | |#| | | |#",
		"############ ####"}

	currentGame := game.Deserialize(rawGame)

	//WHEN
	bestMoves := ai.GetNextMovesScores(currentGame, true)

	//THEN

	downRight := bestMoves["down:right"]
	if downRight != ai.WinningScore {
		t.Errorf("The winning move should be 'down:right' %v", bestMoves)
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
