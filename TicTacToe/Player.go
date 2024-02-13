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
