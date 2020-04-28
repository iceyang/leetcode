package main

func jump(nums []int) int {
	length := len(nums)
	if length <= 0 {
		return 0
	}
	steps := 0
	position := 0
	curMaxPosition := 0
	for i := 0; i < length-1; i++ {
		if position < nums[i]+i {
			position = nums[i] + i
		}
		if i == curMaxPosition {
			curMaxPosition = position
			steps += 1
		}
	}
	return steps
}
