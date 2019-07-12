package main

// 直接在原数组上做处理
func dp(grid [][]int, m int, n int) int {
	if grid[0][0] == 1 {
		return 0
	}
	grid[0][0] = 1
	for i := 1; i < n; i++ {
		if grid[i][0] == 1 || grid[i-1][0] == 0 {
			grid[i][0] = 0
		} else {
			grid[i][0] = 1
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				continue
			}
			if i == 0 {
				grid[i][j] = grid[i][j-1]
				continue
			}
			grid[i][j] = grid[i][j-1] + grid[i-1][j]
		}
	}
	return grid[n-1][m-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	if n == 0 {
		return 0
	}
	m := len(obstacleGrid[0])
	if m == 0 {
		return 0
	}
	return dp(obstacleGrid, m, n)
}
