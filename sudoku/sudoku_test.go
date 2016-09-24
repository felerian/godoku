package sudoku

import (
	"testing"
)

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
