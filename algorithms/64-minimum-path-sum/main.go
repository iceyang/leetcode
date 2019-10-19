package main

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m <= 0 {
		return 0
	}
	n := len(grid[0])
	if n <= 0 {
		return 0
	}
	result := make([][]int, m)
	for i, _ := range result {
		result[i] = make([]int, n)
	}
	result[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		result[i][0] = grid[i][0] + result[i-1][0]
	}
	for j := 1; j < n; j++ {
		result[0][j] = grid[0][j] + result[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			v := grid[i][j]
			result[i][j] = min(result[i-1][j], result[i][j-1]) + v
		}
	}

	return result[m-1][n-1]
}
