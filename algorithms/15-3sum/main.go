package main

import (
	"sort"
)

var result [][]int

func contain(current []int) bool {
	if len(result) == 0 {
		return false
	}
	for _, nums := range result {
		if current[0] == nums[0] &&
			current[1] == nums[1] &&
			current[2] == nums[2] {
			return true
		}
	}
	return false
}

func find(nums []int, current []int, remainTimes int) {
	if remainTimes == 0 {
		num := 0
		for _, v := range current {
			num += v
		}
		if num != 0 {
			return
		}
		sort.Ints(current)
		if !contain(current) {
			result = append(result, current)
		}
	}
	length := len(nums)
	if length <= 0 {
		return
	}
	for i := 0; i < length; i++ {
		find(nums[i+1:], append(current, nums[i]), remainTimes-1)
	}
}

func threeSum1(nums []int) [][]int {
	result = [][]int{}
	find(nums, []int{}, 3)
	return result
}

func threeSum(nums []int) [][]int {
	result = [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		for j := len(nums) - 1; j > i && nums[j] >= 0; j-- {
			sum := nums[i] + nums[j]
			if sum <= 0 {
				for j2 := j - 1; j2 > i && nums[j2] >= -sum; j2-- {
					if sum+nums[j2] == 0 {
						result = append(result, []int{nums[i], nums[j2], nums[j]})
						break
					}
				}
			} else {
				for i2 := i + 1; nums[i2] <= -sum; i2++ {
					if sum+nums[i2] == 0 {
						result = append(result, []int{nums[i], nums[i2], nums[j]})
						break
					}
				}
			}
			for j > i && nums[j] == nums[j-1] {
				j--
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return result
}
