package main

import (
	"fmt"
	"testing"
)

func reverse(nums []int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func nextPermutation(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	i := length - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := length - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)

	nums2 := []int{3, 2, 1}
	nextPermutation(nums2)
	fmt.Println(nums2)

	nums3 := []int{1, 1, 5}
	nextPermutation(nums3)
	fmt.Println(nums3)

	nums4 := []int{1, 3, 2}
	nextPermutation(nums4)
	fmt.Println(nums4)

	nums5 := []int{}
	nextPermutation(nums5)
	fmt.Println(nums5)

	nextPermutation(nil)

	nums6 := []int{1, 5, 8, 4, 7, 6, 5, 3, 1}
	nextPermutation(nums6)
	fmt.Println(nums6)
}
