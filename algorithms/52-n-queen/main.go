package main

import (
	"fmt"
)

func printQueen(chessboard []int) {
	n := len(chessboard)
	for row := 1; row < n; row++ {
		for column := 1; column < n; column++ {
			if chessboard[row] == column {
				fmt.Print("Q")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func check(chessboard []int, row int, column int) bool {
	for i := 1; i < row; i++ {
		if chessboard[i] == column || column-chessboard[i] == row-i || -(column-chessboard[i]) == row-i {
			return false
		}
	}
	return true
}

func totalNQueens(n int) int {
	count := 0
	chessboard := make([]int, n+1)
	for i := 1; i <= n; {
		j := chessboard[i] + 1
		for j <= n {
			if check(chessboard, i, j) {
				chessboard[i] = j
				break
			}
			j++
		}
		if j > n {
			chessboard[i] = 0
			i--
			if i == 0 {
				return count
			}
		} else if i == n {
			printQueen(chessboard)
			count++
		} else {
			i++
		}
	}
	return count
}
