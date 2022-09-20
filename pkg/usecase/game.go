package usecase

import (
	"power4/pkg/domain"
	"power4/pkg/interfaces"
)

type Game interface {
	Play(intrface interfaces.Interface)
}

func NewGame(grid domain.Grid, player1 domain.Player, player2 domain.Player) Game {
	return &game{
		grid:    grid,
		players: []domain.Player{player1, player2},
	}
}

type game struct {
	grid               domain.Grid
	players            []domain.Player
	currentPlayerIndex int
	currentPlayer      domain.Player
}

func (g *game) Play(intrface interfaces.Interface) {
	for {
		currentPlayer := g.nextPlayer()
		correctPlay := false

		intrface.ShowGrid(g.grid)
		for !correctPlay {
			if err := currentPlayer.Play(intrface.AskColumnTo, g.grid); err != nil {
				intrface.Slap(currentPlayer, err)
				intrface.ShowGrid(g.grid)
			} else {
				correctPlay = true
			}
		}
		if currentPlayer.Wins(g.grid) {
			intrface.ShowGrid(g.grid)
			intrface.Congratulate(currentPlayer)
			return
		}
		if g.grid.IsFull() {
			intrface.ShowGrid(g.grid)
			intrface.NoWinner()
			return
		}
	}
}

func (g *game) nextPlayer() domain.Player {
	player := g.players[g.currentPlayerIndex]
	g.currentPlayerIndex = (g.currentPlayerIndex + 1) % 2

	return player
}
