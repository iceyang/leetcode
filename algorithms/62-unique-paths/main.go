package main

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
	return r[n-1][m-1]
}

func dp2(m int, n int) int {
	r := make([]int, m)
	for i, _ := range r {
		r[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			r[j] += r[j-1]
		}
	}
	return r[m-1]
}

func recursion(m, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return recursion(m, n-1) + recursion(m-1, n)
}

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	return dp2(m, n)
}
