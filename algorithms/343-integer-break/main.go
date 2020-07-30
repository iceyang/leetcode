package main

func integerBreak(n int) int {
	result := make([]int, n+1)
	result[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			left := j
			if result[j] > left {
				left = result[j]
			}
			right := i - j
			if result[i-j] > right {
				right = result[i-j]
			}
			value := left * right
			if result[i] < value {
				result[i] = value
			}
		}
	}
	return result[n]
}
