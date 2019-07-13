package main

func numTrees(n int) int {
	r := make([]int, n+1)
	r[0] = 1
	r[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// 左子树为r[j-1]
			// 右子树为r[i-j]
			// 两者的组合（笛卡尔积）就是r[i]当前以j为根的组合数量
			r[i] += r[j-1] * r[i-j]
		}
	}
	return r[n]
}
