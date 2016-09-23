[![Build Status](https://travis-ci.org/felerian/godoku.svg?branch=master)](https://travis-ci.org/felerian/godoku)
[![Coverage Status](https://coveralls.io/repos/github/felerian/godoku/badge.svg?branch=master)](https://coveralls.io/github/felerian/godoku?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/felerian/godoku)](https://goreportcard.com/report/github.com/felerian/godoku)
[![GoDoc](https://godoc.org/github.com/felerian/godoku?status.png)](https://godoc.org/github.com/felerian/godoku)

Godoku
======

A simple sudoku solver, written in Go.

Getting started
---------------

```sh
# fetch "godoku":
go get github.com/felerian/godoku

# change to godoku directory:
cd $GOPATH/src/github.com/felerian/godoku

# run godoku server:
go run godoku.go

# run tests with coverage:
go test -cover ./...
```

To solve a sudoku, provide it as JSON in a POST request:

```sh
curl -X POST localhost:8080/solve -H "Content-Type: application/json" -d @- <<EOF
[
    [0, 0, 0, 0, 0, 3, 0, 8, 0],
    [3, 0, 8, 1, 0, 0, 0, 0, 9],
    [0, 9, 0, 6, 2, 0, 0, 0, 4],
    [0, 0, 9, 0, 0, 7, 0, 2, 0],
    [8, 0, 0, 0, 0, 0, 1, 0, 0],
    [0, 0, 0, 0, 0, 2, 0, 0, 3],
    [0, 0, 0, 7, 0, 6, 8, 0, 0],
    [2, 0, 5, 4, 8, 0, 3, 0, 0],
    [7, 0, 3, 0, 5, 0, 0, 1, 0]
]
EOF
```

The resulting solutions (100 solutions max) are returned as follows:

```json
{"solutions":
    [
        [
            [6, 5, 2, 9, 4, 3, 7, 8, 1],
            [3, 4, 8, 1, 7, 5, 2, 6, 9],
            [1, 9, 7, 6, 2, 8, 5, 3, 4],
            [4, 3, 9, 5, 1, 7, 6, 2, 8],
            [8, 2, 6, 3, 9, 4, 1, 7, 5],
            [5, 7, 1, 8, 6, 2, 9, 4, 3],
            [9, 1, 4, 7, 3, 6, 8, 5, 2],
            [2, 6, 5, 4, 8, 1, 3, 9, 7],
            [7, 8, 3, 2, 5, 9, 4, 1, 6]
        ], [
            [6, 7, 2, 9, 4, 3, 5, 8, 1],
            [3, 4, 8, 1, 7, 5, 2, 6, 9],
            [5, 9, 1, 6, 2, 8, 7, 3, 4],
            [4, 3, 9, 5, 1, 7, 6, 2, 8],
            [8, 2, 6, 3, 9, 4, 1, 7, 5],
            [1, 5, 7, 8, 6, 2, 9, 4, 3],
            [9, 1, 4, 7, 3, 6, 8, 5, 2],
            [2, 6, 5, 4, 8, 1, 3, 9, 7],
            [7, 8, 3, 2, 5, 9, 4, 1, 6]
        ], [
            [6, 4, 2, 9, 7, 3, 5, 8, 1],
            [3, 7, 8, 1, 4, 5, 2, 6, 9],
            [5, 9, 1, 6, 2, 8, 7, 3, 4],
            [4, 3, 9, 5, 1, 7, 6, 2, 8],
            [8, 2, 6, 3, 9, 4, 1, 7, 5],
            [1, 5, 7, 8, 6, 2, 9, 4, 3],
            [9, 1, 4, 7, 3, 6, 8, 5, 2],
            [2, 6, 5, 4, 8, 1, 3, 9, 7],
            [7, 8, 3, 2, 5, 9, 4, 1, 6]
        ]
    ]
}
```
