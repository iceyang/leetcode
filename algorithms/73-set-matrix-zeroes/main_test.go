package main

import (
	"fmt"
	"testing"
)

func printMatrix(matrix [][]int) {
	for _, v := range matrix {
		fmt.Println(v)
	}
}

func TestSetZeroes(t *testing.T) {
	// matrix := [][]int{
	// 	[]int{1, 1, 1},
	// 	[]int{1, 0, 1},
	// 	[]int{1, 1, 1},
	// }
	// fmt.Println("Before:")
	// printMatrix(matrix)

	// setZeroes(matrix)

	// fmt.Println("After:")
	// printMatrix(matrix)

	// fmt.Println("------------------")

	// matrix2 := [][]int{
	// 	[]int{1, 1, 2, 0},
	// 	[]int{3, 4, 5, 2},
	// 	[]int{1, 3, 1, 5},
	// }
	// fmt.Println("Before:")
	// printMatrix(matrix2)

	// setZeroes(matrix2)

	// fmt.Println("After:")
	// printMatrix(matrix2)

	fmt.Println("------------------")

	matrix3 := [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}
	fmt.Println("Before:")
	printMatrix(matrix3)

	setZeroes(matrix3)

	fmt.Println("After:")
	printMatrix(matrix3)
}
