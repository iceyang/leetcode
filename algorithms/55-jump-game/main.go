package main

func canJump(nums []int) bool {
	maxDistance := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		if i > maxDistance {
			return false
		}
		step := nums[i]
		if maxDistance < i+step {
			maxDistance = i + step
		}
	}
	return true
}
