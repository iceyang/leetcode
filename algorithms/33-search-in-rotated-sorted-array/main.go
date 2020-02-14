package main

func search(nums []int, target int) int {
	length := len(nums)
	if length <= 0 {
		return -1
	}
	left, right, mid := 0, length-1, length/2
	for left <= right {
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && nums[mid] >= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[right] >= target && nums[mid] <= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		mid = (left + right) / 2
	}
	return -1
}
