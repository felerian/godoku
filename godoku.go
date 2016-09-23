/*
Package godoku provides a server for solving sudokus via HTTP.
*/
package main

import (
	"encoding/json"
	"github.com/felerian/godoku/sudoku"
	"log"
	"net/http"
	"time"
)

const maxNrOfSolutions int = 100

type response struct {
	Solutions []sudoku.Sudoku `json:"solutions"`
}

func handleSolve(rw http.ResponseWriter, req *http.Request) {
	start := time.Now()

	s := sudoku.Sudoku{}
	if decErr := json.NewDecoder(req.Body).Decode(&s); decErr != nil {
		log.Println(decErr)
		rw.WriteHeader(400)
		return
	}
	solutions := s.Solve(maxNrOfSolutions)
	response := response{
		Solutions: solutions,
	}
	json.NewEncoder(rw).Encode(response)

	duration := time.Since(start).Nanoseconds() / 1000000
	log.Printf("Responded with %d solutions after %d ms.", len(solutions), duration)
}

func main() {
	log.Println("Godoku server started.")
	http.HandleFunc("/solve", handleSolve)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
