[![Build Status](https://travis-ci.org/felerian/godoku.svg?branch=master)](https://travis-ci.org/felerian/godoku)

Godoku
======

A simple sudoku solver, written in Go.

How to use
----------

```
// fetch "godoku":
go get github.com/felerian/godoku

// change to godoku directory:
cd $GOPATH/src/github.com/felerian/godoku

// run godoku server:
go run godoku.go

// run tests with coverage:
go test -cover ./...
```

To solve a sudoku, provide it as JSON in a POST request:

```
curl -kv -X POST localhost:8080/solve -H "Content-Type: application/json" -d @- <<EOF
[[0, 0, 0, 0, 0, 3, 0, 8, 0],
 [3, 0, 8, 1, 0, 0, 0, 0, 9],
 [0, 9, 0, 6, 2, 0, 0, 0, 4],
 [0, 0, 9, 0, 0, 7, 0, 2, 0],
 [8, 0, 0, 0, 0, 0, 1, 0, 0],
 [0, 0, 0, 0, 0, 2, 0, 0, 3],
 [0, 0, 0, 7, 0, 6, 8, 0, 0],
 [2, 0, 0, 4, 8, 0, 3, 0, 0],
 [0, 0, 3, 0, 5, 0, 0, 1, 0]]
EOF
```

The resulting solutions (100 solutions max.) are returned as a JSON array.
