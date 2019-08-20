package main

import (
	"sort"
)

func contains(result [][]int, newArr []int) bool {
	for _, arr := range result {
		if len(arr) != len(newArr) {
			continue
		}
		exists := true
		for i, val := range arr {
			if newArr[i] != val {
				exists = false
				break
			}
		}
		if exists {
			return true
		}
	}
	return false
}

func backtrack(candidates [][]int, target int) [][]int {
	numbers := candidates[target]
	var result [][]int
	for _, num := range numbers {
		if num-target == 0 {
			result = append(result, []int{num})
		} else {
			for _, res := range backtrack(candidates, target-num) {
				newArr := append(res, num)
				sort.Ints(newArr)
				if !contains(result, newArr) {
					result = append(result, newArr)
				}
			}
		}
	}
	return result
}

func combinationSum(candidates []int, target int) [][]int {
	if target == 0 {
		return [][]int{[]int{}}
	}
	result := make([][]int, target+1)
	result[0] = []int{0}
	for i := 1; i <= target; i++ {
		for _, num := range candidates {
			if i-num < 0 {
				continue
			}
			if len(result[i-num]) > 0 {
				result[i] = append(result[i], num)
			}
		}
	}
	return backtrack(result, target)
}
