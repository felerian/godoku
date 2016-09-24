/*
Package solver provides means to solve sudokus.
*/
package solver

import (
	"github.com/felerian/godoku/digitset"
	"github.com/felerian/godoku/sudoku"
)

// Solver for a sudoku
type Solver [9][9]digitset.DigitSet

// Solve a sudoku
func Solve(sudoku sudoku.Sudoku, maxSolutions int) []sudoku.Sudoku {
	solver := prepare(sudoku)
	count := 0
	return recursiveSolve(solver, maxSolutions, &count)
}

// prepare creates a solver from a sudoku
func prepare(sudoku sudoku.Sudoku) Solver {
	solver := Solver{}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if value := sudoku[r][c]; value > 0 && value < 10 {
				solver[r][c] = digitset.Single(value)
			} else {
				solver[r][c] = digitset.All()
			}
		}
	}
	return solver
}

func recursiveSolve(s Solver, maxSolutions int, count *int) []sudoku.Sudoku {
	s.simplify()
	if !s.valid() {
		return []sudoku.Sudoku{}
	}
	if s.solved() {
		*count++
		return []sudoku.Sudoku{s.flatten()}
	}
	solutions := []sudoku.Sudoku{}
	r, c := s.findMinimumChoices()
	choices := s[r][c]
	for i := uint(0); i < 9; i++ {
		if choices.Contains(i) {
			s[r][c] = digitset.Single(i)
			if *count < maxSolutions {
				solutions = append(solutions, recursiveSolve(s, maxSolutions, count)...)
			} else {
				return solutions
			}
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

// valid returns true, if this solver is valid (contains no contradictions)
func (s *Solver) valid() bool {
	v := validByGroup(s, row)
	v = v && validByGroup(s, col)
	v = v && validByGroup(s, block)
	return v
}

// validByGroup validates this solver along a row, column or block
func validByGroup(s *Solver, accessor func(*Solver, uint) func(uint) *digitset.DigitSet) bool {
	for g := uint(0); g < 9; g++ {
		set := digitset.Empty()
		group := accessor(s, g)
		for i := uint(0); i < 9; i++ {
			if value, err := group(i).Value(); err == nil {
				if set.Contains(value) {
					return false
				}
				set.Add(value)
			}
		}
	}
	return true
}

// flatten turns a solver back into a sudoku
func (s *Solver) flatten() sudoku.Sudoku {
	sudoku := sudoku.Sudoku{}
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
	r := 3*(f/3) + i/3
	c := 3*(f%3) + i%3
	return r, c
}

// simplify this solver in-place without resorting to guesses
func (s *Solver) simplify() {
	changed := true
	for changed {
		changed = simplifyByGroup(s, row)
		changed = changed || simplifyByGroup(s, col)
		changed = changed || simplifyByGroup(s, block)
	}
}

// simplifyByGroup simplifies this solver in-place along a row, column or block
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
