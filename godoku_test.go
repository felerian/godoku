package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSolveSudokuWithManySolutions(t *testing.T) {
	// given
	payload := []byte(`[
		[0, 0, 0, 0, 0, 3, 0, 8, 0],
		[3, 0, 8, 1, 0, 0, 0, 0, 9],
		[0, 9, 0, 6, 2, 0, 0, 0, 4],
		[0, 0, 9, 0, 0, 7, 0, 2, 0],
		[8, 0, 0, 0, 0, 0, 1, 0, 0],
		[0, 0, 0, 0, 0, 2, 0, 0, 3],
		[0, 0, 0, 7, 0, 6, 8, 0, 0],
		[2, 0, 0, 4, 8, 0, 3, 0, 0],
		[0, 0, 3, 0, 5, 0, 0, 1, 0]
	]`)
	req, _ := http.NewRequest("POST", "localhost:8080", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rw := httptest.NewRecorder()
	expectedCount := 92
	// when
	handleSolve(rw, req)
	// then
	response := Response{}
	json.Unmarshal([]byte(rw.Body.String()), &response)
	count := len(response.Solutions)
	if rw.Code != 200 {
		t.Errorf("expected: status OK (200), actual: %d", rw.Code)
	}
	if count != expectedCount {
		t.Errorf("expected: %d solutions, actual: %d", expectedCount, count)
	}
}

func TestSolveSudokuWithTooManySolutions(t *testing.T) {
	// given
	payload := []byte(`[
		[0, 0, 0, 0, 0, 0, 0, 8, 0],
		[0, 0, 0, 0, 0, 0, 0, 0, 9],
		[0, 9, 0, 6, 2, 0, 0, 0, 4],
		[0, 0, 9, 0, 0, 7, 0, 2, 0],
		[8, 0, 0, 0, 0, 0, 1, 0, 0],
		[0, 0, 0, 0, 0, 2, 0, 0, 3],
		[0, 0, 0, 7, 0, 6, 8, 0, 0],
		[2, 0, 0, 4, 8, 0, 3, 0, 0],
		[0, 0, 3, 0, 5, 0, 0, 1, 0]
	]`)
	req, _ := http.NewRequest("POST", "localhost:8080", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rw := httptest.NewRecorder()
	expectedCount := MAX_SOLUTIONS
	// when
	handleSolve(rw, req)
	// then
	response := Response{}
	json.Unmarshal([]byte(rw.Body.String()), &response)
	count := len(response.Solutions)
	if rw.Code != 200 {
		t.Errorf("expected: status OK (200), actual: %d", rw.Code)
	}
	if count != expectedCount {
		t.Errorf("expected: %d solutions, actual: %d", expectedCount, count)
	}
}

func TestFailSolvingForMalformedInput(t *testing.T) {
	// given
	payload := []byte(`gibberish`)
	req, _ := http.NewRequest("POST", "localhost:8080", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rw := httptest.NewRecorder()
	// when
	handleSolve(rw, req)
	// then
	if rw.Code != 400 {
		t.Errorf("expected: status Bad Request (400), actual: %d", rw.Code)
	}
}
