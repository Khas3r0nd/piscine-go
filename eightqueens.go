package piscine

import (
	"github.com/01-edu/z01"
)

// function to check if a position is safe to place a queen
func IsSafe(queenNum int, rowPos int, pos []int) bool {
	// check if the position is in the same row, diagonal, or column as any of the previously placed queens
	for i := 0; i < queenNum; i++ {
		t := pos[i]
		if t == rowPos || t == rowPos-(queenNum-i) || t == rowPos+(queenNum-i) {
			// if it is, return false
			return false
		}
	}
	// if the position is safe, return true
	return true
}

// function that generates all possible solutions to the eight queens puzzle
func solve(queenNum int, pos []int) {
	// if all 8 queens have been placed on the board
	if queenNum == 8 {
		// print the solution
		for _, v := range pos {
			z01.PrintRune(rune(v + 49))
		}
		z01.PrintRune(10)
	} else {
		// try placing a queen on each square of the board
		for i := 0; i < 8; i++ {
			// if the position is safe
			if IsSafe(queenNum, i, pos) {
				// place the queen on the board
				pos[queenNum] = i
				// move on to the next queen
				solve(queenNum+1, pos)
			}
		}
	}
}

// function to start solving the eight queens puzzle
func EightQueens() {
	// initializes the board
	pos := []int{0, 0, 0, 0, 0, 0, 0, 0}
	// starts solving the puzzle
	solve(0, pos)
}
