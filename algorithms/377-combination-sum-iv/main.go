package main

import (
	"sort"
)

/**
 * 动态规划解答问题
 */
func combinationSum4(candidates []int, target int) int {
	sort.Ints(candidates)
	result := make([]int, target+1)
	result[0] = 1
	for i := 1; i <= target; i++ {
		result[i] = 0
		for _, num := range candidates {
			if i-num < 0 {
				continue
			} else if i-num == 0 {
				result[i] += 1
			} else if result[i-num] > 0 {
				result[i] += result[i-num]
			}
		}
	}
	return result[target]
}
