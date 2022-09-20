package main

import (
	"power4/pkg/domain"
	"power4/pkg/infrastructure/local"
	"power4/pkg/interfaces"
	"power4/pkg/interfaces/console"
	"power4/pkg/usecase"
)

var (
	grid     domain.Grid
	player1  domain.Player
	player2  domain.Player
	game     usecase.Game
	intrface interfaces.Interface
)

func init() {
	grid = local.NewGrid()
	player1 = local.NewPlayer('x')
	player2 = local.NewPlayer('o')
	game = usecase.NewGame(grid, player1, player2)
	intrface = console.NewConsole()
}

func main() {
	game.Play(intrface)
}
