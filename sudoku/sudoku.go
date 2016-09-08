package sudoku

import (
	"github.com/felerian/godoku/digitset"
	"strconv"
)

type Template [9][9]uint

type Sudoku [9][9]digitset.DigitSet

func New() Sudoku {
	return Sudoku{}
}

func (s *Sudoku) String() string {
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
			set := (*s)[r][c]
			if value, err := set.Value(); err == nil {
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

func Init(t Template) Sudoku {
	sudoku := New()
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if value := t[r][c]; value > 0 && value < 10 {
				sudoku[r][c] = digitset.Single(value)
			} else {
				sudoku[r][c] = digitset.All()
			}
		}
	}
	return sudoku
}

// Row returns a function yielding the elements of the r'th row
func Row(sudoku *Sudoku, r uint) func(uint) *digitset.DigitSet {
	return func(c uint) *digitset.DigitSet {
		return &(*sudoku)[r][c]
	}
}

// Col returns a function yielding the elements of the c'th column
func Col(sudoku *Sudoku, c uint) func(uint) *digitset.DigitSet {
	return func(r uint) *digitset.DigitSet {
		return &(*sudoku)[r][c]
	}
}

/*
// Field returns a function yielding the elements of the f'th field
func Field(sudoku *Sudoku, f uint) func(uint) *digitset.DigitSet {
	return func(i uint) *digitset.DigitSet {
		return &(*sudoku)[row][col]
	}
}
*/

func (s *Sudoku) Simplify() {
	var changed bool = true
	for changed {
		changed = SimplifyByGroup(s, Row)
		changed = changed || SimplifyByGroup(s, Col)
		//changed = changed || SimplifyByGroup(s, Field)
	}
}

func SimplifyByGroup(s *Sudoku, accessor func(*Sudoku, uint) func(uint) *digitset.DigitSet) bool {
	var changed bool
	for g := uint(0); g < 9; g++ {
		group := accessor(s, g)
		for i := uint(0); i < 9; i++ {
			if vi, ei := group(i).Value(); ei == nil {
				for j := uint(0); j < 9; j++ {
					if _, ej := group(j).Value(); ej != nil && group(j).Contains(vi) {
						group(j).Remove(vi)
						changed = true
					}
				}
			}
		}
	}
	return changed
}
