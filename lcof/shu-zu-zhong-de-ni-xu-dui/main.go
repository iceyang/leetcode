package main

import "fmt"

func reversePairs(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	result := make([]int, length)
	result[length-1] = 0

	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j <= length-1; j++ {
			if nums[i] < nums[j] {
				continue
			}
			result[i] += result[j]
			if nums[i] > nums[j] {
				result[i] += 1
			}
		}
	}
	fmt.Println(result)

	count := 0
	for _, v := range result {
		count += v
	}

	return count
}
