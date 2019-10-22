package main

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	if n == 0 {
		return []int{}
	}
	i := 0
	j := 0
	count := 0
	length := m * n
	res := [length]int{}
	// for count < length {
	// 	res[count] = matrix[i][j]
	// 	count++
	// 	if i < m {
	// 	}
	// }
}
