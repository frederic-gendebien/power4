package domain

type Player interface {
	GetSymbol() rune
	Play(ask func(Player) (int, error), grid Grid) error
	Wins(grid Grid) bool
}
