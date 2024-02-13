package TicTacToe

import "fmt"

type Player struct {
	name           string
	symbol         rune
	penaltyCounter int
}

func CreatePlayer(name string, symbol rune) Player {
	return Player{
		name,
		symbol,
		0,
	}
}

func (player Player) GetSymbol() rune {
	return player.symbol
}

func (player Player) PrintPlayer() {
	fmt.Printf("Name: %s Symbol: %c\n", player.name, player.symbol)
}

func (player *Player) IncrementPenalty() {
	player.penaltyCounter = 1 + player.penaltyCounter
}

func PrintPlayersInformation(player1,player2 Player){
	fmt.Println("Player 1:")
    fmt.Println("Name:", player1.name)
    fmt.Println("Symbol:", string(player1.symbol))
    fmt.Println("Penalty Counter:", player1.penaltyCounter)

    fmt.Println("\nPlayer 2:")
    fmt.Println("Name:", player2.name)
    fmt.Println("Symbol:", string(player2.symbol))
    fmt.Println("Penalty Counter:", player2.penaltyCounter)

    fmt.Println("\n" + "------------------------------------------")
}
