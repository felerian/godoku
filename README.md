[![Build Status](https://travis-ci.org/felerian/godoku.svg?branch=master)](https://travis-ci.org/felerian/godoku)

Godoku
======

A simple sudoku solver, written in Go.

Start the server:

```
go run godoku.go
```

Call with a sudoku to solve:

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
