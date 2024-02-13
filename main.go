package main

import (
	"fmt"
	TTT "main/TicTacToe"
)

func printMenu() {
	fmt.Println("Main Menu:")
	fmt.Println("1. Play Game")
	fmt.Println("2. Instructions")
	fmt.Println("3. Exit")
	fmt.Println("Please enter your choice: [1/2/3]")
}

func printInstructions() {
    fmt.Println("Instructions:")
    fmt.Println("1. The game is played on a 3x3 grid.")
    fmt.Println("2. Two players take turns marking an empty cell with their respective symbols.")
    fmt.Println("3. The first player to get three of their symbols in a row (horizontally, vertically, or diagonally) wins.")
    fmt.Println("4. If all cells are filled and no player has three in a row, the game is a draw.")
    fmt.Println("5. To make a move, enter the row and column numbers corresponding to the cell.")
    fmt.Println("   For example, to place your symbol in the top-right corner, enter '1 3'.")
    fmt.Println("6. If you try to occupy an already filled cell, it will be considered a penalty.")
    fmt.Println("   Consecutively doing so for 3 times will result in a defeat.")
    fmt.Println("7. Have fun playing!")
}

func initiateGame(){
	var player1 = TTT.CreatePlayer("P1", 'O')
	var player2 = TTT.CreatePlayer("P2", 'X')
	var game = TTT.CreateGame(player1, player2)
	game.StartGame()
	fmt.Println("Final:")
	TTT.PrintGame(game)
}

func main() {
	fmt.Println("Welcome to Tic Tac Toe!")
	mainLoop:
	for {
		printMenu()
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			initiateGame()
		case 2:
			printInstructions()
		case 3:
			break mainLoop
		default:
			fmt.Println("Wrong input!!! Please choose from menu!!!")
		}
	}

}
