package domain

type Grid interface {
	Process(func([6][7]rune))
	DropCoin(column int, symbol rune) error
	IsFull() bool
	Has4InARow(interesting rune) bool
}
