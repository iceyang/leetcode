package main

import "math"

func findKElementWithTwoSortedArrays(nums1 *[]int, nums2 *[]int, i int, j int, k int) int {
	if i >= len((*nums1)) {
		return (*nums2)[j+k-1]
	}
	if j >= len((*nums2)) {
		return (*nums1)[i+k-1]
	}
	if k == 1 {
		if (*nums1)[i] <= (*nums2)[j] {
			return (*nums1)[i]
		} else {
			return (*nums2)[j]
		}
	}
	num1 := math.MaxInt64
	num2 := math.MaxInt64
	if i+k/2-1 < len((*nums1)) {
		num1 = (*nums1)[i+k/2-1]
	}
	if j+k/2-1 < len((*nums2)) {
		num2 = (*nums2)[j+k/2-1]
	}
	if num1 < num2 {
		return findKElementWithTwoSortedArrays(nums1, nums2, i+k/2, j, k-k/2)
	} else {
		return findKElementWithTwoSortedArrays(nums1, nums2, i, j+k/2, k-k/2)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	result := findKElementWithTwoSortedArrays(&nums1, &nums2, 0, 0, (length+1)/2)
	if length%2 == 0 {
		return float64(result+findKElementWithTwoSortedArrays(&nums1, &nums2, 0, 0, (length+2)/2)) / 2
	}
	return float64(result)
}

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
