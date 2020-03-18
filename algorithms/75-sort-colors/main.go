package main

func sortColors(nums []int) {
	left, cur, right := 0, 0, len(nums)-1
	if right <= 0 {
		return
	}
	for cur <= right {
		if nums[cur] == 0 {
			nums[cur], nums[left] = nums[left], nums[cur]
			left++
			cur++
		} else if nums[cur] == 2 {
			nums[cur], nums[right] = nums[right], nums[cur]
			right--
		} else {
			cur++
		}
	}
}
