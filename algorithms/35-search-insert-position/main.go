package main

func searchInsert(nums []int, target int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == target || nums[i] > target {
			return i
		}
	}
	return length
}
