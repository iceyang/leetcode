package main

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	f := make([]int, n+1)
	f[1] = 1
	f[2] = 2
	i := 3
	for ; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
