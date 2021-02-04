package tests

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"tipsy/ai"
	"tipsy/game"
)


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
	moves := ai.GetNextMovesScores(currentGame, 2, true)

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
	moves := ai.GetNextMovesScores(currentGame, 2, true)

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
	moves := ai.GetNextMovesScores(game, 2, true)

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
	moves := ai.GetNextMovesScores(game, 2, true)

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
	bestMoves := ai.GetNextMovesScores(currentGame, 2, true)

	//THEN

	downRight := bestMoves["down:right"]
	if downRight != ai.WinningScore {
		t.Errorf("The winning move should be 'down:right' %v", bestMoves)
	}
}

//one blue one red out => should not win
func TestMoveToDownRightShouldntBeTheBestMove(t *testing.T) {

	//GIVEN
	rawGame := []string{"blue",
		"#### ############",
		"#| | | |#|B| |b|#",
		"#| |#| |B| |#|R| ",
		"#| |R|#|R|#|B| |#",
		"#|#|R| |R|B| |#|#",
		"#| | |#| |#| |B|#",
		" | |#| | | |#| |#",
		"#| | | |#| |r| |#",
		"############ ####"}

	currentGame := game.Deserialize(rawGame)

	//WHEN
	bestMoves := ai.GetNextMovesScores(currentGame, 2, true)

	//THEN

	downRight := bestMoves["down:right"]
	if downRight == ai.WinningScore {
		t.Errorf("The winning move should be 'down:right' %v", bestMoves)
	}
}

//two blue out in 2 moves
func TestMoveToDownRightShouldBeTheBestMoveAsItPushOutLastBluePucks(t *testing.T) {

	//GIVEN
	rawGame := []string{"blue",
		"#### ############",
		"#| | | |#| | |b|#",
		"#| |#| |B| |#|R| ",
		"#| |R|#|R|#|B| |#",
		"#|#|R| |R|B| |#|#",
		"#| | |#| |#| |B|#",
		" | |#| | |r|#| |#",
		"#| | | |#| |b| |#",
		"############ ####"}

	currentGame := game.Deserialize(rawGame)

	//WHEN
	bestMoves := ai.GetNextMovesScores(currentGame, 1, true)

	//THEN

	downRight := bestMoves["down:right"]
	if downRight != ai.WinningScore {
		t.Errorf("The winning move should be 'down:right' %v", bestMoves)
	}
}

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
