package main

import (
	"sort"
)

/**
 * 计算result中是否包含了newArr
 */
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

var result [][]int

/**
 * 回溯解答问题（从大往小）
 */
func backtrack2(candidates *[]int, remain int, curIndex int, nums *[]int) {
	length := len(*candidates)
	for i := curIndex; i < length; i++ {
		num := (*candidates)[i]
		if num > remain {
			// 剪枝
			break
		}
		if num == remain {
			newArr := append(*nums, num)
			if !contains(result, newArr) {
				result = append(result, newArr)
			}
			break
		}
		if i == length-1 {
			// 说明后面已经没有数字了
			break
		}
		nums2 := make([]int, len(*nums))
		copy(nums2, *nums)
		nums2 = append(nums2, num)
		backtrack2(candidates, remain-num, i+1, &nums2)
		// TODO
	}
}

/**
 * 回溯解答问题（从小往大）
 */
func backtrack(candidates *[]int, target int, sum int, curIndex int, nums *[]int) {
	length := len(*candidates)
	for i := curIndex; i < length; i++ {
		num := (*candidates)[i]
		if num+sum == target {
			result = append(result, append(*nums, num))
			break
		}
		if num+sum > target {
			break
		}
		if i != curIndex && (*candidates)[i] == (*candidates)[i-1] {
			continue
		}
		nums2 := make([]int, len(*nums))
		copy(nums2, *nums)
		nums2 = append(nums2, num)
		backtrack(candidates, target, num+sum, i+1, &nums2)

	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	if target == 0 {
		return [][]int{[]int{}}
	}
	result = [][]int{}
	// backtrack2(&candidates, target, 0, &[]int{})
	backtrack(&candidates, target, 0, 0, &[]int{})
	return result
}
