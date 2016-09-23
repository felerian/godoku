package sudoku

import (
	"github.com/felerian/godoku/digitset"
	"testing"
)

func TestRows(t *testing.T) {
	// given
	solver := Solver{}
	// when
	group := row(&solver, 0)
	for i := uint(0); i < 9; i++ {
		group(i).Add(i + 1)
	}
	// then
	for c := uint(0); c < 9; c++ {
		if value, err := solver[0][c].Value(); value != c+1 || err != nil {
			t.Errorf("expected: %d, actual: %d", value, c)
		}
	}
}

func TestCols(t *testing.T) {
	// given
	solver := Solver{}
	// when
	group := col(&solver, 0)
	for i := uint(0); i < 9; i++ {
		group(i).Add(i + 1)
	}
	// then
	for r := uint(0); r < 9; r++ {
		if value, err := solver[r][0].Value(); value != r+1 || err != nil {
			t.Errorf("expected: %d, actual: %d", value, r)
		}
	}
}

func TestFieldToCoords(t *testing.T) {
	// given
	var f uint = 2
	var i uint = 1
	// when
	r, c := blocksToCoords(f, i)
	// then
	if r != 0 || c != 7 {
		t.Errorf("expected: (0, 7), actual: (%d, %d)", r, c)
	}
}

func TestSimplifyRow(t *testing.T) {
	// given
	solver := Solver{}
	for c := uint(0); c < 9; c++ {
		if c == 4 {
			solver[0][c] = digitset.All()
		} else {
			solver[0][c] = digitset.Single(c + 1)
		}
	}
	if value, err := solver[0][4].Value(); err == nil {
		t.Errorf("Should be undetermined (%+v)", value)
	}
	// when
	changed := simplifyByGroup(&solver, row)
	// then
	if !changed {
		t.Error("Should have been changed.")
	}
	if value, err := solver[0][4].Value(); value != 5 || err != nil {
		t.Errorf("Should have been simplified to 5 (%+v)", value)
	}
}

func TestSimplifyCol(t *testing.T) {
	// given
	solver := Solver{}
	for r := uint(0); r < 9; r++ {
		if r == 6 {
			solver[r][0] = digitset.All()
		} else {
			solver[r][0] = digitset.Single(r + 1)
		}
	}
	if value, err := solver[6][0].Value(); err == nil {
		t.Errorf("Should be undetermined (%+v)", value)
	}
	// when
	changed := simplifyByGroup(&solver, col)
	// then
	if !changed {
		t.Error("Should have been changed.")
	}
	if value, err := solver[6][0].Value(); value != 7 || err != nil {
		t.Errorf("Should have been simplified to 7 (%+v)", value)
	}
}

func TestSimplifyField(t *testing.T) {
	// given
	solver := Solver{}
	for i := uint(0); i < 9; i++ {
		r, c := blocksToCoords(0, i)
		if i == 8 {
			solver[r][c] = digitset.All()
		} else {
			solver[r][c] = digitset.Single(i + 1)
		}
	}
	if value, err := solver[2][2].Value(); err == nil {
		t.Errorf("Should be undetermined (%+v)", value)
	}
	// when
	changed := simplifyByGroup(&solver, block)
	// then
	if !changed {
		t.Error("Should have been changed.")
	}
	if value, err := solver[2][2].Value(); value != 9 || err != nil {
		t.Errorf("Should have been simplified to 9 (%+v)", value)
	}
}

func TestInit(t *testing.T) {
	// given
	var sudoku Sudoku = [9][9]uint{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	// when
	solver := sudoku.prepare()
	// then
	if value, err := solver[0][0].Value(); value != 1 || err != nil {
		t.Errorf("expected: 1, actual: %d", value)
	}
	if _, err := solver[1][1].Value(); err == nil {
		t.Error("Should be unknown.")
	}
}

func TestSimplify(t *testing.T) {
	// given
	var sudoku Sudoku = [9][9]uint{
		{7, 2, 0, 0, 0, 3, 0, 8, 1},
		{3, 0, 8, 1, 0, 0, 0, 6, 9},
		{0, 9, 0, 6, 2, 8, 0, 0, 4},
		{6, 0, 9, 5, 0, 7, 0, 2, 0},
		{8, 5, 2, 0, 9, 0, 1, 0, 0},
		{0, 0, 0, 0, 6, 2, 9, 5, 3},
		{0, 1, 5, 7, 0, 6, 8, 0, 0},
		{2, 6, 0, 4, 8, 0, 3, 0, 0},
		{0, 0, 3, 0, 5, 0, 6, 1, 7},
	}
	solver := sudoku.prepare()
	// when
	solver.simplify()
	// then
	if !solver.solved() {
		t.Error("Should be solved.")
	}
}

func TestShouldBeValid(t *testing.T) {
	// given
	var sudoku Sudoku = [9][9]uint{
		{7, 2, 0, 0, 0, 3, 0, 8, 1},
		{3, 0, 8, 1, 0, 0, 0, 6, 9},
		{0, 9, 0, 6, 2, 8, 0, 0, 4},
		{6, 0, 9, 5, 0, 7, 0, 2, 0},
		{8, 5, 2, 0, 9, 0, 1, 0, 0},
		{0, 0, 0, 0, 6, 2, 9, 5, 3},
		{0, 1, 5, 7, 0, 6, 8, 0, 0},
		{2, 6, 0, 4, 8, 0, 3, 0, 0},
		{0, 0, 3, 0, 5, 0, 6, 1, 7},
	}
	solver := sudoku.prepare()
	// when
	v := solver.valid()
	// then
	if !v {
		t.Error("Should be valid.")
	}
}

func TestShouldBeInvalidForInvalidRow(t *testing.T) {
	// given
	var sudoku Sudoku = [9][9]uint{
		{7, 2, 0, 0, 0, 3, 0, 8, 2},
		{3, 0, 8, 1, 0, 0, 0, 6, 9},
		{0, 9, 0, 6, 2, 8, 0, 0, 4},
		{6, 0, 9, 5, 0, 7, 0, 2, 0},
		{8, 5, 2, 0, 9, 0, 1, 0, 0},
		{0, 0, 0, 0, 6, 2, 9, 5, 3},
		{0, 1, 5, 7, 0, 6, 8, 0, 0},
		{2, 6, 0, 4, 8, 0, 3, 0, 0},
		{0, 0, 3, 0, 5, 0, 6, 1, 7},
	}
	solver := sudoku.prepare()
	// when
	v := solver.valid()
	// then
	if v {
		t.Error("Should be invalid.")
	}
}

func TestShouldBeInvalidForInvalidColumn(t *testing.T) {
	// given
	var sudoku Sudoku = [9][9]uint{
		{7, 2, 0, 0, 0, 3, 0, 8, 1},
		{3, 0, 8, 1, 0, 0, 0, 6, 9},
		{0, 9, 0, 6, 2, 8, 0, 0, 4},
		{6, 0, 9, 5, 0, 7, 0, 2, 0},
		{8, 5, 2, 0, 9, 0, 1, 0, 0},
		{0, 0, 0, 0, 6, 2, 9, 5, 3},
		{0, 1, 5, 7, 0, 6, 8, 0, 0},
		{7, 6, 0, 4, 8, 0, 3, 0, 0},
		{0, 0, 3, 0, 5, 0, 6, 1, 7},
	}
	solver := sudoku.prepare()
	// when
	v := solver.valid()
	// then
	if v {
		t.Error("Should be invalid.")
	}
}

func TestShouldBeInvalidForInvalidBlock(t *testing.T) {
	// given
	var sudoku Sudoku = [9][9]uint{
		{7, 7, 0, 0, 0, 3, 0, 8, 1},
		{3, 0, 8, 1, 0, 0, 0, 6, 9},
		{0, 9, 0, 6, 2, 8, 0, 0, 4},
		{6, 0, 9, 5, 0, 7, 0, 2, 0},
		{8, 5, 2, 0, 9, 0, 1, 0, 0},
		{0, 0, 0, 0, 6, 2, 9, 5, 3},
		{0, 1, 5, 7, 0, 6, 8, 0, 0},
		{2, 6, 0, 4, 8, 0, 3, 0, 0},
		{0, 0, 3, 0, 5, 0, 6, 1, 7},
	}
	solver := sudoku.prepare()
	// when
	v := solver.valid()
	// then
	if v {
		t.Error("Should be invalid.")
	}
}

func TestFindSingleSolution(t *testing.T) {
	// given
	var sudoku Sudoku = [9][9]uint{
		{7, 2, 0, 0, 0, 3, 0, 8, 1},
		{3, 0, 8, 1, 0, 0, 0, 6, 9},
		{0, 9, 0, 6, 2, 8, 0, 0, 4},
		{6, 0, 9, 5, 0, 7, 0, 2, 0},
		{8, 5, 2, 0, 9, 0, 1, 0, 0},
		{0, 0, 0, 0, 6, 2, 9, 5, 3},
		{0, 1, 5, 7, 0, 6, 8, 0, 0},
		{2, 6, 0, 4, 8, 0, 3, 0, 0},
		{0, 0, 3, 0, 5, 0, 6, 1, 7},
	}
	maxSolutions := 3
	// when
	solutions := sudoku.Solve(maxSolutions)
	// then
	if len(solutions) != 1 {
		t.Errorf("expected: 1 solutions, actual: %d", len(solutions))
	}
}

func TestFindThreeSolutions(t *testing.T) {
	// given
	var sudoku Sudoku = [9][9]uint{
		{0, 0, 0, 0, 0, 3, 0, 8, 0},
		{3, 0, 8, 1, 0, 0, 0, 6, 9},
		{0, 9, 0, 6, 2, 8, 0, 0, 4},
		{6, 0, 9, 5, 0, 7, 0, 2, 0},
		{8, 5, 2, 0, 9, 0, 1, 0, 0},
		{0, 0, 0, 0, 6, 2, 9, 5, 3},
		{0, 1, 5, 7, 0, 6, 8, 0, 0},
		{2, 6, 0, 4, 8, 0, 3, 0, 0},
		{0, 0, 3, 0, 5, 0, 6, 1, 7},
	}
	maxSolutions := 3
	// when
	solutions := sudoku.Solve(maxSolutions)
	// then
	if len(solutions) != maxSolutions {
		t.Errorf("expected: %d solutions, actual: %d", maxSolutions, len(solutions))
	}
}

func TestPrintSudoku(t *testing.T) {
	// given
	sudoku := Sudoku{}
	sudoku[4][4] = 1
	// when
	stringRepr := sudoku.String()
	// then
	if !(stringRepr == "\n"+
		"  _ _ _  _ _ _  _ _ _\n"+
		"  _ _ _  _ _ _  _ _ _\n"+
		"  _ _ _  _ _ _  _ _ _\n"+
		"\n"+
		"  _ _ _  _ _ _  _ _ _\n"+
		"  _ _ _  _ 1 _  _ _ _\n"+
		"  _ _ _  _ _ _  _ _ _\n"+
		"\n"+
		"  _ _ _  _ _ _  _ _ _\n"+
		"  _ _ _  _ _ _  _ _ _\n"+
		"  _ _ _  _ _ _  _ _ _\n\n") {
		t.Error("Incorrect string representation.")
	}
}
