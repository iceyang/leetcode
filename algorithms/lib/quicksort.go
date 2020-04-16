package lib

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	guard := nums[0]
	i := 0
	j := len(nums) - 1
	for i < j {
		for nums[j] >= guard && i < j {
			j--
		}
		nums[i] = nums[j]
		for nums[i] < guard && i < j {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = guard
	QuickSort(nums[:i])
	QuickSort(nums[i+1:])
	return nums
}
