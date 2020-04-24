package main

// 解法一
// 暴力，肯定超时
func reversePairs1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	result := make([]int, length)
	result[length-1] = 0

	for i := length - 2; i >= 0; i-- {
		result[i] = result[i+1]
		for j := i + 1; j <= length-1; j++ {
			if nums[i] > nums[j] {
				result[i] += 1
			}
		}
	}

	return result[0]
}

// 归并排序解
func mergeSort(nums []int, start int, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	left := mergeSort(nums, start, mid)
	right := mergeSort(nums, mid+1, end)
	count := 0
	tmp := make([]int, end-start+1)
	i, j, cur := start, mid+1, 0
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp[cur] = nums[i]
			cur++
			count += j - (mid + 1)
			i++
		} else {
			tmp[cur] = nums[j]
			cur++
			j++
		}
	}
	remainI := mid - i + 1
	if remainI > 0 {
		count += (end - mid) * remainI
	}
	for ; i <= mid; i++ {
		tmp[cur] = nums[i]
		cur++
	}
	for ; j <= end; j++ {
		tmp[cur] = nums[j]
		cur++
	}
	copy(nums[start:end+1], tmp)
	return count + left + right
}

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}
