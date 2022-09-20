package local

import "fmt"

func NewGrid() *grid {
	var matrix [6][7]rune
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			matrix[i][j] = ' '
		}
	}

	return &grid{
		Matrix:    matrix,
		available: 6 * 7,
	}
}

type grid struct {
	Matrix    [6][7]rune
	available int
}

func (g *grid) Process(processing func([6][7]rune)) {
	processing(g.Matrix)
}

func (g *grid) DropCoin(column int, symbol rune) error {
	if column < 0 || column > 6 {
		return fmt.Errorf("you entered wrong column, dumbass")
	}

	freeRow := -1
	for _, row := range g.Matrix {
		if row[column] != ' ' {
			break
		}
		freeRow++
	}
	if freeRow == -1 {
		return fmt.Errorf("the column is full, are you blind?")
	}

	g.Matrix[freeRow][column] = symbol
	g.available--

	return nil
}

func (g *grid) IsFull() bool {
	return g.available == 0
}

func (g *grid) Has4InARow(interesting rune) bool {
	verticals := []int{0, 0, 0, 0, 0, 0, 0}
	dDiagonals := []int{0, 0, 0, 0, 0, 0}
	aDiagonals := []int{0, 0, 0, 0, 0, 0}

	update := func(values *[]int, index int, update func(int) int) {
		if index >= 0 && index < len(*values) {
			(*values)[index] = update((*values)[index])
		}
	}

	has4 := func(values []int, index int) bool {
		return index >= 0 && index < len(values) && values[index] == 4
	}

	for rowIndex, row := range g.Matrix {
		horizontal := 0
		for columnIndex, symbol := range row {
			dDiagIndex := columnIndex - rowIndex + 2
			aDiagIndex := columnIndex + rowIndex - 3
			if symbol == interesting {
				horizontal++
				update(&verticals, columnIndex, increment())
				update(&dDiagonals, dDiagIndex, increment())
				update(&aDiagonals, aDiagIndex, increment())
			} else {
				horizontal = 0
				update(&verticals, columnIndex, zero())
				update(&dDiagonals, dDiagIndex, zero())
				update(&aDiagonals, aDiagIndex, zero())
			}
			if horizontal == 4 || has4(verticals, columnIndex) || has4(dDiagonals, dDiagIndex) || has4(aDiagonals, aDiagIndex) {
				return true
			}
		}
	}

	return false
}

func increment() func(int) int {
	return func(value int) int {
		return value + 1
	}
}

func zero() func(int) int {
	return func(value int) int {
		return 0
	}
}
