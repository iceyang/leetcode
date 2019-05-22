package main

func RemoveElement(nums []int, val int) int {
	if nums == nil {
		return 0
	}
	i := 0
	j := len(nums)
	for i < j {
		if nums[i] == val {
			j--
			nums[i] = nums[j]
			continue
		}
		i++
	}
	return j
}
