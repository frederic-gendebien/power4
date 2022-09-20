package console

import (
	"bufio"
	"fmt"
	"os"
	"power4/pkg/domain"
	"strconv"
)

func NewConsole() Console {
	return Console{}
}

type Console struct {
}

func (c Console) ShowGrid(grid domain.Grid) {
	printColumns := func() {
		for i := 1; i < 8; i++ {
			fmt.Printf("   %d", i)
		}
		fmt.Println()
	}

	fmt.Println()
	printColumns()
	grid.Process(func(matrix [6][7]rune) {
		for _, line := range matrix {
			for _, symbole := range line {
				fmt.Print(" | ", string(symbole))
			}
			fmt.Println(" | ")
		}
	})
	printColumns()
}

func (c Console) AskColumnTo(player domain.Player) (int, error) {
	normalizeColumn := func(col rune, _ int, err error) (int, error) {
		if err != nil {
			return -1, err
		}

		result, err := strconv.Atoi(string(col))
		if err != nil {
			return -1, fmt.Errorf("enter a number you dumbass!")
		}

		if result < 1 || result > 7 {
			return -1, fmt.Errorf("enter a number between 1 and 7, you dumbass!")
		}

		return result - 1, nil
	}

	fmt.Println()
	fmt.Printf("player '%s' enter column: ", string(player.GetSymbol()))
	reader := bufio.NewReader(os.Stdin)
	col, err := normalizeColumn(reader.ReadRune())
	if err != nil {
		return -1, err
	}

	return col, nil
}

func (c Console) Congratulate(player domain.Player) {
	fmt.Println()
	fmt.Printf("congratulations player '%s' you win!", string(player.GetSymbol()))
	fmt.Println()
}

func (c Console) Slap(player domain.Player, err error) {
	fmt.Println()
	fmt.Printf("player '%s' %s", string(player.GetSymbol()), err.Error())
	fmt.Println()
}

func (c Console) NoWinner() {
	fmt.Println()
	fmt.Println("nobody wins, losers!")
	fmt.Println()
}
