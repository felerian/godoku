/*
Package godoku provides a server for solving sudokus via HTTP.
*/
package main

import (
	"encoding/json"
	"github.com/felerian/godoku/sudoku"
	"log"
	"net/http"
)

const maxNrOfSolutions int = 100

type response struct {
	Solutions []sudoku.Sudoku `json:"solutions"`
}

func handleSolve(rw http.ResponseWriter, req *http.Request) {
	s := sudoku.Sudoku{}
	if decErr := json.NewDecoder(req.Body).Decode(&s); decErr != nil {
		log.Println(decErr)
		rw.WriteHeader(400)
		return
	}
	response := response{
		Solutions: s.Solve(maxNrOfSolutions),
	}
	json.NewEncoder(rw).Encode(response)
}

func main() {
	log.Println("Godoku server started.")
	http.HandleFunc("/solve", handleSolve)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
