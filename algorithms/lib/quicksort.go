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
		for nums[i] <= guard && i < j {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	nums[0], nums[i] = nums[i], nums[0]
	QuickSort(nums[:i])
	QuickSort(nums[i+1:])
	return nums
}
