package main

import (
	"fmt"
	"github.com/felerian/godoku/digitset"
	"github.com/felerian/godoku/sudoku"
)

func main() {
	// Sudoku board
	sudoku := sudoku.New()
	sudoku[2][3].Add(4)
	sudoku.Simplify()
	fmt.Printf(sudoku.String())

	set := digitset.Empty()
	set.Add(1)
	set.Add(3)
	set.Remove(1)
	set.Remove(3)

	fmt.Printf("%d\n", set)
}
