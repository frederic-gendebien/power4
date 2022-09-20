package interfaces

import (
	"power4/pkg/domain"
)

type Interface interface {
	ShowGrid(grid domain.Grid)
	AskColumnTo(player domain.Player) (int, error)
	Congratulate(player domain.Player)
	Slap(player domain.Player, err error)
	NoWinner()
}
