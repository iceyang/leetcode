package main

/**
 * 动态规划解法
 */
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSteps(n int) int {
	steps := make([]int, n+1)
	for i := 2; i <= n; i++ {
		steps[i] = i
		for j := 1; j <= i/2; j++ {
			if i%j != 0 {
				continue
			}
			steps[i] = min(steps[i], steps[j]+i/j)
		}
	}
	return steps[n]
}

/** 利用n跟质数的关系求解
 * 假设n为质数，那么它只能有n步，不能通过其它数字复制得到
 * 如果不是质数，则可以由复制得到
 */
// func minSteps(n int) int {
// 	step := 0
// 	for i := 2; i <= n; i++ {
// 		for n%i == 0 {
// 			step += i
// 			n /= i
// 		}
// 	}
// 	return step
// }
