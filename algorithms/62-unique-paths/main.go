package main

import "fmt"

func dp(m int, n int) int {
	r := make([][]int, n)
	for i, _ := range r {
		r[i] = make([]int, m)
		r[i][0] = 1
	}

	for i := 0; i < m; i++ {
		r[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			r[i][j] = r[i][j-1] + r[i-1][j]
		}
	}

	for _, v := range r {
		fmt.Println(v)
	}
	return r[n-1][m-1]
}

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	return dp(m, n)
}
