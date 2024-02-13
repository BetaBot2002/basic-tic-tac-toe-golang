package TicTacToe

import (
	"fmt"
	"math/rand"
)

var defaultSymbols = []rune{'O', 'X'}

type Game struct {
	Board         [][]int
	Player1       Player
	Player2       Player
	CurrentPlayer Player
	Winner        Player
	IsDrawn       bool
	OccupiedCells int
}

func GetRandomSymbol() rune {
	RandomIndex := rand.Intn(2)
	return defaultSymbols[RandomIndex]
}

func CreateGame(player1, player2 Player) Game {
	board := [][]int{
		{-1, -1, -1},
		{-1, -1, -1},
		{-1, -1, -1},
	}

	return Game{
		board,
		player1,
		player2,
		player1,
		Player{},
		false,
		0,
	}
}

func PrintBoard(board [][]int) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
	fmt.Println("-------------------------------------")
}

func PrintGame(game Game) {
	// game.Player1.PrintPlayer()
	// game.Player2.PrintPlayer()
	PrintBoard(game.Board)
}

func (game Game) CheckForWinner(row, col int) bool {
	board := game.Board
	symbol := int(game.CurrentPlayer.GetSymbol())
	rowCheck := (board[row][0] == symbol) && (board[row][1] == symbol) && (board[row][2] == symbol)
	colCheck := (board[0][col] == symbol) && (board[1][col] == symbol) && (board[2][col] == symbol)
	leftDiagonalCheck := (board[0][0] == symbol) && (board[1][1] == symbol) && (board[2][2] == symbol)
	rightDiagonalCheck := (board[0][2] == symbol) && (board[1][1] == symbol) && (board[2][0] == symbol)

	return rowCheck || colCheck || leftDiagonalCheck || rightDiagonalCheck
}

func (game *Game) getCurrentPlayerPenalty() int {
	if game.CurrentPlayer.name != game.Player1.name {
		return game.Player2.penaltyCounter
	}
	return game.Player1.penaltyCounter
}

func (game *Game) incrementCurrentPlayerPenalty() {
	if game.CurrentPlayer.name != game.Player1.name {
		game.Player2.IncrementPenalty()
		return
	}
	game.Player1.IncrementPenalty()
}

func (game Game) getNextPlayer() Player {
	if game.CurrentPlayer.name == game.Player1.name {
		return game.Player2
	}
	return game.Player1
}

func (game Game) StartGame() {
	for !game.IsDrawn {
		fmt.Println("Current Turn: ", game.CurrentPlayer.name)
		var row, col int
		fmt.Println("Enter the row number:")
		fmt.Scan(&row)
		row = row % 3
		fmt.Println("Enter the column number:")
		fmt.Scan(&col)
		col = col % 3
		if game.Board[row][col] == -1 {
			game.Board[row][col] = int(game.CurrentPlayer.GetSymbol())
			game.OccupiedCells++
		} else {
			fmt.Println("The Cell you mentioned is already occupied.\nThis is against the game rule.\nViolation of this rule for 3 times will result in a win for your opponent.")
			game.incrementCurrentPlayerPenalty()
			fmt.Println("Current Penalty: ", game.getCurrentPlayerPenalty())
			fmt.Println("You have ", 3-game.getCurrentPlayerPenalty(), " more chances.")
			if game.getCurrentPlayerPenalty() == 3 {
				game.Winner = game.getNextPlayer()
				break
			}
		}
		PrintGame(game)
		if game.CheckForWinner(row, col) {
			game.Winner = game.CurrentPlayer
			break
		}

		game.CurrentPlayer = game.getNextPlayer()
		// fmt.Println("Turn Debug: ", game.CurrentPlayer)

		game.IsDrawn = game.OccupiedCells >= 9
	}

	fmt.Println("Game Over!!!")

	if game.Winner != (Player{}) {
		fmt.Println("The Winner is:")
		game.Winner.PrintPlayer()
	}else if game.IsDrawn{
		fmt.Println("Game Drawn!!!")
	}
}
