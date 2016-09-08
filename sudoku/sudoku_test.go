package sudoku

import (
	"github.com/felerian/godoku/digitset"
	"testing"
)

func TestRows(t *testing.T) {
	// given
	sudoku := New()
	// when
	row := Row(&sudoku, 0)
	for i := uint(0); i < 9; i++ {
		row(i).Add(i + 1)
	}
	// then
	for c := uint(0); c < 9; c++ {
		if value, err := sudoku[0][c].Value(); value != c+1 || err != nil {
			t.Errorf("expected: %d, actual: %d", value, c)
		}
	}
}

func TestCols(t *testing.T) {
	// given
	sudoku := New()
	// when
	col := Col(&sudoku, 0)
	for i := uint(0); i < 9; i++ {
		col(i).Add(i + 1)
	}
	// then
	for r := uint(0); r < 9; r++ {
		if value, err := sudoku[r][0].Value(); value != r+1 || err != nil {
			t.Errorf("expected: %d, actual: %d", value, r)
		}
	}
}

func TestSimplifyRow(t *testing.T) {
	// given
	sudoku := New()
	for c := uint(0); c < 9; c++ {
		if c == 4 {
			sudoku[0][c] = digitset.All()
		} else {
			sudoku[0][c] = digitset.Single(c + 1)
		}
	}
	if value, err := sudoku[0][4].Value(); err == nil {
		t.Errorf("Should be undetermined (%+v)", value)
	}
	// when
	changed := SimplifyByGroup(&sudoku, Row)
	// then
	if !changed {
		t.Error("Should have been changed.")
	}
	if value, err := sudoku[0][4].Value(); value != 5 || err != nil {
		t.Errorf("Should have been simplified to 5 (%+v)", value)
	}
}

func TestSimplifyCol(t *testing.T) {
	// given
	sudoku := New()
	for r := uint(0); r < 9; r++ {
		if r == 6 {
			sudoku[r][0] = digitset.All()
		} else {
			sudoku[r][0] = digitset.Single(r + 1)
		}
	}
	if value, err := sudoku[6][0].Value(); err == nil {
		t.Errorf("Should be undetermined (%+v)", value)
	}
	// when
	changed := SimplifyByGroup(&sudoku, Col)
	// then
	if !changed {
		t.Error("Should have been changed.")
	}
	if value, err := sudoku[6][0].Value(); value != 7 || err != nil {
		t.Errorf("Should have been simplified to 7 (%+v)", value)
	}
}
