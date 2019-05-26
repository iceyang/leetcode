package main

func TwoSum(nums []int, target int) []int {
	res := map[int]int{}
	res[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		remain := target - nums[i]
		remainIndex, exists := res[remain]
		if exists {
			return []int{remainIndex, i}
		}
		res[nums[i]] = i
	}
	return []int{0, len(nums) - 1}
}

/**
 * violent version.
 */
func TwoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, len(nums) - 1}
}
