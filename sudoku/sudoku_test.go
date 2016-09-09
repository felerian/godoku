package sudoku

import (
	"github.com/felerian/godoku/digitset"
	"testing"
)

func TestRows(t *testing.T) {
	// given
	sudoku := Sudoku{}
	// when
	group := row(&sudoku, 0)
	for i := uint(0); i < 9; i++ {
		group(i).Add(i + 1)
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
	sudoku := Sudoku{}
	// when
	group := col(&sudoku, 0)
	for i := uint(0); i < 9; i++ {
		group(i).Add(i + 1)
	}
	// then
	for r := uint(0); r < 9; r++ {
		if value, err := sudoku[r][0].Value(); value != r+1 || err != nil {
			t.Errorf("expected: %d, actual: %d", value, r)
		}
	}
}

func TestFieldToCoords(t *testing.T) {
	// given
	var f uint = 2
	var i uint = 1
	// when
	r, c := fieldsToCoords(f, i)
	// then
	if r != 0 || c != 7 {
		t.Errorf("expected: (0, 7), actual: (%d, %d)", r, c)
	}
}

func TestSimplifyRow(t *testing.T) {
	// given
	sudoku := Sudoku{}
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
	changed := SimplifyByGroup(&sudoku, row)
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
	sudoku := Sudoku{}
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
	changed := SimplifyByGroup(&sudoku, col)
	// then
	if !changed {
		t.Error("Should have been changed.")
	}
	if value, err := sudoku[6][0].Value(); value != 7 || err != nil {
		t.Errorf("Should have been simplified to 7 (%+v)", value)
	}
}

func TestSimplifyField(t *testing.T) {
	// given
	sudoku := Sudoku{}
	for i := uint(0); i < 9; i++ {
		r, c := fieldsToCoords(0, i)
		if i == 8 {
			sudoku[r][c] = digitset.All()
		} else {
			sudoku[r][c] = digitset.Single(i + 1)
		}
	}
	if value, err := sudoku[2][2].Value(); err == nil {
		t.Errorf("Should be undetermined (%+v)", value)
	}
	// when
	changed := SimplifyByGroup(&sudoku, field)
	// then
	if !changed {
		t.Error("Should have been changed.")
	}
	if value, err := sudoku[2][2].Value(); value != 9 || err != nil {
		t.Errorf("Should have been simplified to 9 (%+v)", value)
	}
}

func TestInit(t *testing.T) {
	// given
	var template Template = [9][9]uint{
		[9]uint{1, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]uint{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	// when
	sudoku := Init(template)
	// then
	if value, err := sudoku[0][0].Value(); value != 1 || err != nil {
		t.Errorf("expected: 1, actual: %d", value)
	}
	if _, err := sudoku[1][1].Value(); err == nil {
		t.Error("Should be unknown.")
	}
}

func TestSimplify(t *testing.T) {
	// given
	var template Template = [9][9]uint{
		[9]uint{7, 2, 0, 0, 0, 3, 0, 8, 1},
		[9]uint{3, 0, 8, 1, 0, 0, 0, 6, 9},
		[9]uint{0, 9, 0, 6, 2, 8, 0, 0, 4},
		[9]uint{6, 0, 9, 5, 0, 7, 0, 2, 0},
		[9]uint{8, 5, 2, 0, 9, 0, 1, 0, 0},
		[9]uint{0, 0, 0, 0, 6, 2, 9, 5, 3},
		[9]uint{0, 1, 5, 7, 0, 6, 8, 0, 0},
		[9]uint{2, 6, 0, 4, 8, 0, 3, 0, 0},
		[9]uint{0, 0, 3, 0, 5, 0, 6, 1, 7},
	}
	sudoku := Init(template)
	// when
	sudoku.Simplify()
	// then
	if !sudoku.Solved() {
		t.Error("Should be solved.")
	}
}
