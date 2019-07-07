package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * 动态规划，自上向下
 * d[i][j] = min{ d[i-1][j-1], d[i-1][j] } + triangle[i][j]
 */
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	res := make([][]int, len(triangle))
	for i := range res {
		res[i] = make([]int, i+1)
	}
	res[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			v := triangle[i][j]
			if j == 0 {
				res[i][j] = res[i-1][j] + v
				continue
			}
			if j == len(triangle[i])-1 {
				res[i][j] = res[i-1][j-1] + v
				continue
			}
			res[i][j] = min(res[i-1][j-1], res[i-1][j]) + v
		}
	}
	r := res[len(res)-1][0]
	for _, v := range res[len(res)-1] {
		r = min(r, v)
	}
	return r
}

/**
 * 动态规划，自底向上
 * d[j] = min{ d[j], d[j+1] } + triangle[i][j]
 * 只需要 O(n) 额外空间
 */
func minimumTotal2(triangle [][]int) int {
	row := len(triangle)
	if row == 0 {
		return 0
	}
	res := make([]int, row)
	for i, v := range triangle[row-1] {
		res[i] = v
	}
	for i := row - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			v := triangle[i][j]
			res[j] = min(res[j], res[j+1]) + v
		}
	}
	return res[0]
}
