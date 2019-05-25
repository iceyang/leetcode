package main

func TwoSum(nums []int, target int) []int {
	return []int{0, len(nums) - 1}
}

/**
 * violent version.
 */
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, len(nums) - 1}
}
