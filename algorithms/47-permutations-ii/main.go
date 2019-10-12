package main

import (
	"fmt"
	"sort"
)

var result [][]int

func backtrack(nums []int, cur int) {
	length := len(nums)
	if length == cur {
		fmt.Println(nums)
		result = append(result, append([]int{}, nums...))
		return
	}
	used := map[int]bool{}
	for i := cur; i < length; i++ {
		num := nums[i]
		if _, ok := used[num]; ok {
			continue
		}
		used[num] = true
		nums[cur], nums[i] = nums[i], nums[cur]
		backtrack(nums, cur+1)
		nums[cur], nums[i] = nums[i], nums[cur]
	}
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result = make([][]int, 0, 10)
	backtrack(nums, 0)
	return result
}
