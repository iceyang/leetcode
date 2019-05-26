package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	index := TwoSum(nums, target)
	if nums[index[0]]+nums[index[1]] != target {
		t.Errorf("expected %d but got result %d with index %+v. (input %+v)", target, nums[index[0]]+nums[index[1]], index, nums)
	}
}
