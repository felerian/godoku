package main

import (
	"encoding/json"
	"github.com/felerian/godoku/sudoku"
	"log"
	"net/http"
)

type Response struct {
	Solutions []sudoku.Sudoku
}

func handleSolve(rw http.ResponseWriter, req *http.Request) {
	s := sudoku.Sudoku{}
	if decErr := json.NewDecoder(req.Body).Decode(&s); decErr != nil {
		log.Println(decErr)
		rw.WriteHeader(400)
		return
	}
	response := Response{
		Solutions: s.Solve(),
	}
	json.NewEncoder(rw).Encode(response)
}

func main() {
	log.Println("Godoku server started.")
	http.HandleFunc("/solve", handleSolve)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
