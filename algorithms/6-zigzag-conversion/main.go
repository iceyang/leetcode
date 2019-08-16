package main

import "fmt"

func convert(s string, numRows int) string {
	res := make([][]rune, numRows)
	for i, _ := range res {
		res[i] = make([]rune, 20)
	}
	// count := (numRows * (numRows - 1))
	for i, letter := range s {
		// i % count
		row := i % numRows
		column := i / numRows
		res[row][column] = letter
	}
	for _, r := range res {
		fmt.Println(string(r))
	}
	return ""
}
