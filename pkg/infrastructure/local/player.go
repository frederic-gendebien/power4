package local

import (
	"fmt"
	"power4/pkg/domain"
)

func NewPlayer(symbol rune) Player {
	return Player{
		Symbol: symbol,
	}
}

type Player struct {
	Symbol rune
}

func (p Player) GetSymbol() rune {
	return p.Symbol
}

func (p Player) Play(ask func(domain.Player) (int, error), grid domain.Grid) error {
	column, err := ask(p)
	if err != nil {
		return fmt.Errorf("what did you do? %w", err)
	}

	return grid.DropCoin(column, p.Symbol)
}

func (p Player) Wins(grid domain.Grid) bool {
	return grid.Has4InARow(p.Symbol)
}
