package main

// 时间复杂度未达到要求的版本
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	length1 := len(nums1)
	length2 := len(nums2)
	length := length1 + length2
	target := length / 2
	index1, index2 := 0, 0
	var cur int
	for target > 0 {
		if index1 >= length1 {
			// 数组1耗完
			cur = nums2[index2]
			index2++
		} else if index2 >= length2 {
			// 数组2耗完
			cur = nums1[index1]
			index1++
		} else if nums1[index1] <= nums2[index2] {
			cur = nums1[index1]
			index1++
		} else {
			cur = nums2[index2]
			index2++
		}
		target--
	}
	var result float64
	if index1 >= length1 || (index2 < length2 && nums1[index1] > nums2[index2]) {
		result = float64(nums2[index2])
	} else {
		result = float64(nums1[index1])
	}
	if length%2 == 0 {
		result = (result + float64(cur)) / 2
	}
	return result
}
