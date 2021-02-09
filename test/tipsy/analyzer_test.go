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
	_, moves := ai.GetNextMovesScores(currentGame, 3, true)

	//THEN
	rightLeftWin := moves["right:left"]
	rightDownWin := moves["right:down"]
	rightRightWin := moves["right:right"]
	rightUpWin := moves["right:up"]
	if rightLeftWin != ai.WinningScore {
		t.Errorf("The rightLeftWin move score should be WinningScore %v", moves)
	}
	if rightDownWin != ai.WinningScore {
		t.Errorf("The rightDownWin move score should be WinningScore %v", moves)
	}
	if rightRightWin != ai.WinningScore {
		t.Errorf("The rightRightWin move score should be WinningScore %v", moves)
	}
	if rightUpWin != ai.WinningScore {
		t.Errorf("The rightUpWin move score should be WinningScore %v", moves)
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
	_, moves := ai.GetNextMovesScores(currentGame, 3, true)

	//THEN
	rightLeftWin := moves["right:left"]
	rightDownWin := moves["right:down"]
	rightRightWin := moves["right:right"]
	rightUpWin := moves["right:up"]
	if rightLeftWin != ai.LosingScore {
		t.Errorf("The rightLeftWin move score should be LosingScore %v", moves)
	}
	if rightDownWin != ai.LosingScore {
		t.Errorf("The rightDownWin move score should be LosingScore %v", moves)
	}
	if rightRightWin != ai.LosingScore {
		t.Errorf("The rightRightWin move score should be LosingScore %v", moves)
	}
	if rightUpWin != ai.LosingScore {
		t.Errorf("The rightUpWin move score should be LosingScore %v", moves)
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
	_, moves := ai.GetNextMovesScores(game, 3, true)

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
	bestMove, moves := ai.GetNextMovesScores(game, 3, true)

	//THEN
	if bestMove != "left:down" {
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
	bestMove, bestMoves := ai.GetNextMovesScores(currentGame, 3, true)

	//THEN

	downRight := bestMoves["down:right"]
	if downRight != ai.WinningScore {
		t.Errorf("The 'down:right' move should have winning score %v", bestMoves)
	}
	if bestMove != "down:right" {
		t.Errorf("The winning move should be 'down:right' %v", bestMoves)

	}
}

//one blue one red out => should not win
func TestMoveToDownShouldBeALosingMoveAsItExitRedPuck(t *testing.T) {

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
	_, bestMoves := ai.GetNextMovesScores(currentGame, 3, true)

	//THEN

	downRight := bestMoves["down:right"]
	downLeft := bestMoves["down:left"]
	downUp := bestMoves["down:up"]
	downDown := bestMoves["down:down"]
	if downRight != ai.LosingScore {
		t.Errorf("The 'down:right' move should have LosingScore %v", bestMoves)
	}
	if downLeft != ai.LosingScore {
		t.Errorf("The 'down:left' move should have LosingScore %v", bestMoves)
	}
	if downUp != ai.LosingScore {
		t.Errorf("The 'down:up' move should have LosingScore %v", bestMoves)
	}
	if downDown != ai.LosingScore {
		t.Errorf("The 'down:down' move should have LosingScore %v", bestMoves)
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
	bestMove, bestMoves := ai.GetNextMovesScores(currentGame, 3, true)

	//THEN

	downRight := bestMoves["down:right"]
	if downRight != ai.WinningScore {
		t.Errorf("The 'down:right' move should have the winning score %v", bestMoves)
	}
	if bestMove != "down:right" {
		t.Errorf("The best move should be 'down:right' %v", bestMoves)
	}
}

//black out in one move
func TestMoveToDownRightShouldBeTheBestMoveAsItPushOutTheBlackPuck(t *testing.T) {

	//GIVEN
	rawGame := []string{"blue",
		"#### ############",
		"#| | | |#| | |x|#",
		"#| |#| |B| |#|R| ",
		"#| |R|#|R|#|B| |#",
		"#|#|R| |R|B| |#|#",
		"#| | |#| |#| |B|#",
		" | |#| | | |#| |#",
		"#| | | |#|b|r| |#",
		"############ ####"}

	currentGame := game.Deserialize(rawGame)

	//WHEN
	bestMove, bestMoves := ai.GetNextMovesScores(currentGame, 3, true)

	//THEN

	downRight := bestMoves["down:right"]
	if downRight != ai.WinningScore {
		t.Errorf("The 'down:right' move should have the winning score %v", bestMoves)
	}
	if bestMove != "down:right" {
		t.Errorf("The best move should be 'down:right' %v", bestMoves)
	}
}

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
