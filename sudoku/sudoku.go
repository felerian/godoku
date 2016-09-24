/*
Package sudoku provides a sudoku data structure.
*/
package sudoku

import (
	"strconv"
)

// Sudoku board
type Sudoku [9][9]uint

// String returns a string representation of a sudoku
func (sudoku *Sudoku) String() string {
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
			value := (*sudoku)[r][c]
			if value > 0 && value < 10 {
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
