package sudoku

import (
	"github.com/felerian/godoku/digitset"
	"strconv"
)

type Template [9][9]uint

type Sudoku [9][9]digitset.DigitSet

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
	sudoku := Sudoku{}
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
func row(sudoku *Sudoku, r uint) func(uint) *digitset.DigitSet {
	return func(c uint) *digitset.DigitSet {
		return &(*sudoku)[r][c]
	}
}

// Col returns a function yielding the elements of the c'th column
func col(sudoku *Sudoku, c uint) func(uint) *digitset.DigitSet {
	return func(r uint) *digitset.DigitSet {
		return &(*sudoku)[r][c]
	}
}

// Field returns a function yielding the elements of the f'th field
func field(sudoku *Sudoku, f uint) func(uint) *digitset.DigitSet {
	return func(i uint) *digitset.DigitSet {
		r, c := fieldsToCoords(f, i)
		return &(*sudoku)[r][c]
	}
}

func fieldsToCoords(f uint, i uint) (uint, uint) {
	var r uint = 3*(f/3) + i/3
	var c uint = 3*(f%3) + i%3
	return r, c
}

func (s *Sudoku) Solved() bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if _, err := s[r][c].Value(); err != nil {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) Simplify() {
	var changed bool = true
	for changed {
		changed = SimplifyByGroup(s, row)
		changed = changed || SimplifyByGroup(s, col)
		changed = changed || SimplifyByGroup(s, field)
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
