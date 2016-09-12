package main

import (
	"fmt"
	"github.com/felerian/godoku/sudoku"
)

func main() {
	var sudoku sudoku.Sudoku = [9][9]uint{
		[9]uint{0, 0, 0, 0, 0, 3, 0, 8, 0},
		[9]uint{3, 0, 8, 1, 0, 0, 0, 0, 9},
		[9]uint{0, 9, 0, 6, 2, 0, 0, 0, 4},
		[9]uint{0, 0, 9, 0, 0, 7, 0, 2, 0},
		[9]uint{8, 0, 0, 0, 0, 0, 1, 0, 0},
		[9]uint{0, 0, 0, 0, 0, 2, 0, 0, 3},
		[9]uint{0, 0, 0, 7, 0, 6, 8, 0, 0},
		[9]uint{2, 0, 0, 4, 8, 0, 3, 0, 0},
		[9]uint{0, 0, 3, 0, 5, 0, 0, 1, 0},
	}
	fmt.Println("problem:")
	fmt.Print(sudoku.String())

	solutions := sudoku.Solve()
	nr := len(solutions)

	fmt.Printf("\n%d solutions:\n", nr)
	//for i := 0; i < nr; i++ {
	//	fmt.Print(solutions[i].String())
	//}
}
