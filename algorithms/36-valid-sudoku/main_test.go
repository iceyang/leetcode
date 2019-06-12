package main

import (
	"fmt"
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
	board := make([][]byte, 9)
	arr := []string{
		"..4...63.",
		".........",
		"5......9.",
		"...56....",
		"4.3.....1",
		"...7.....",
		"...5.....",
		".........",
		".........",
	}
	for i, str := range arr {
		board[i] = []byte(str)
	}
	fmt.Println(isValidSudoku(board))
}
