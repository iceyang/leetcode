package main

var result [][]int
var arr []int

func backtrack(nums []int, cur int) {
	length := len(nums)
	if length == cur {
		result = append(result, append([]int{}, nums...))
		return
	}
	for i := cur; i < length; i++ {
		nums[cur], nums[i] = nums[i], nums[cur]
		backtrack(nums, cur+1)
		nums[cur], nums[i] = nums[i], nums[cur]
	}
}

func permute(nums []int) [][]int {
	result = make([][]int, 0, 10)
	arr = make([]int, len(nums))
	backtrack(nums, 0)
	return result
}
