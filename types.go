package main


type Position struct {
	x, y int
}

type Board [][]string

type SolverResult struct {
	word string
	traversedPositions []Position
}