package main

func search(nums *[]int, target int, from int, end int) int {
	if from > end || end < 0 || from > len(*nums)-1 {
		return -1
	}
	i := (from + end) / 2
	if target < (*nums)[i] {
		return search(nums, target, from, i-1)
	} else if target > (*nums)[i] {
		return search(nums, target, i+1, end)
	} else {
		return i
	}
}

func searchRange(nums []int, target int) []int {
	length := len(nums)
	index := search(&nums, target, 0, length-1)
	result := []int{index, index}
	if index == -1 {
		return result
	}
	i := index - 1
	for ; i >= 0; i-- {
		if nums[i] != target {
			break
		}
		result[0] = i
	}
	j := index + 1
	for ; j < length; j++ {
		if nums[j] != target {
			break
		}
		result[1] = j
	}
	return result
}
