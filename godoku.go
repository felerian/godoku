package main

import (
	"fmt"
	"strconv"
)

type Board [9][9]DigitSet

func (b Board) String() string {
	var result string
	for r := 0; r < 9; r++ {
		if r%3 == 0 {
			result += "\n"
		}
		for c := 0; c < 9; c++ {
			result += " "
			if c%3 == 0 {
				result += " "
			}
			ds := b[r][c]
			if value, err := ds.Value(); err == nil {
				result += strconv.Itoa(int(value))
			} else {
				result += "_"
			}
		}
		result += "\n"
	}
	result += "\n"
	return result
}

func (b Board) Recalculate() Board {
	newBoard := Board{}

	// reset unknown fields
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			ds := b[r][c]
			if value, err := ds.Value(); err == nil {
				newBoard[r][c].Add(value)
			} else {
				newBoard[r][c] = All()
			}
		}
	}

	// reduce options by rows
	//...

	// reduce options by columns
	//...

	// reduce options by fields
	//...

	return newBoard
}

func main() {
	// Sudoku board
	b := Board{}
	b[2][3].Add(4)
	b = b.Recalculate()
	fmt.Printf(b.String())

	var ds DigitSet = 0
	ds.Add(1)
	ds.Add(3)
	ds.Remove(1)
	ds.Remove(3)

	fmt.Printf("%d\n", ds)
}
