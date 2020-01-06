package main

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	pre := 0
	result := length
	for i := 1; i < length; i++ {
		if nums[i] == nums[pre] {
			result -= 1
			continue
		}
		pre += 1
		nums[pre] = nums[i]
	}
	return result
}
