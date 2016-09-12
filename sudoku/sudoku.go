package sudoku

import (
	"fmt"
	"github.com/felerian/godoku/digitset"
	"strconv"
)

type Sudoku [9][9]uint

// String returns a string represantation of a sudoku
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
			value := (*s)[r][c]
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

// Solve a sudoku
func (sudoku *Sudoku) Solve() []Sudoku {
	solver := sudoku.prepare()
	return recursiveSolve(solver)
}

// prepare creates a solver from a sudoku
func (s *Sudoku) prepare() Solver {
	solver := Solver{}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if value := s[r][c]; value > 0 && value < 10 {
				solver[r][c] = digitset.Single(value)
			} else {
				solver[r][c] = digitset.All()
			}
		}
	}
	return solver
}

type Solver [9][9]digitset.DigitSet

func recursiveSolve(s Solver) []Sudoku {
	fmt.Print("recursion\n")
	s.simplify()
	if s.solved() {
		return []Sudoku{s.flatten()}
	}
	solutions := []Sudoku{}
	r, c := s.findMinimumChoices()
	choices := s[r][c]
	for i := uint(0); i < 9; i++ {
		if choices.Contains(i) {
			s[r][c] = digitset.Single(i)
			solutions = append(solutions, recursiveSolve(s)...)
		}
	}
	return solutions
}

func (s *Solver) findMinimumChoices() (uint, uint) {
	var row, col uint
	var currentCount uint
	var count uint = 9
	for r := uint(0); r < 9; r++ {
		for c := uint(0); c < 9; c++ {
			currentCount = s[r][c].Count()
			if currentCount > 1 && currentCount < count {
				count = currentCount
				row = r
				col = c
			}
		}
	}
	return row, col
}

// solved returns true, if this solver's work is done
func (s *Solver) solved() bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if _, err := s[r][c].Value(); err != nil {
				return false
			}
		}
	}
	return true
}

// flatten turns a solver back into a sudoku
func (s *Solver) flatten() Sudoku {
	sudoku := Sudoku{}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if value, err := s[r][c].Value(); err != nil {
				sudoku[r][c] = 0
			} else {
				sudoku[r][c] = value
			}
		}
	}
	return sudoku
}

// row returns a function yielding the elements of the r'th row
func row(solver *Solver, r uint) func(uint) *digitset.DigitSet {
	return func(c uint) *digitset.DigitSet {
		return &(*solver)[r][c]
	}
}

// col returns a function yielding the elements of the c'th column
func col(solver *Solver, c uint) func(uint) *digitset.DigitSet {
	return func(r uint) *digitset.DigitSet {
		return &(*solver)[r][c]
	}
}

// block returns a function yielding the elements of the f'th block
func block(solver *Solver, f uint) func(uint) *digitset.DigitSet {
	return func(i uint) *digitset.DigitSet {
		r, c := blocksToCoords(f, i)
		return &(*solver)[r][c]
	}
}

func blocksToCoords(f uint, i uint) (uint, uint) {
	var r uint = 3*(f/3) + i/3
	var c uint = 3*(f%3) + i%3
	return r, c
}

// simplify this solver in-place without resorting to guesses
func (s *Solver) simplify() {
	var changed bool = true
	for changed {
		changed = simplifyByGroup(s, row)
		changed = changed || simplifyByGroup(s, col)
		changed = changed || simplifyByGroup(s, block)
	}
}

// simplifyByGroup this solver in-place along a row, column or block
func simplifyByGroup(s *Solver, accessor func(*Solver, uint) func(uint) *digitset.DigitSet) bool {
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
