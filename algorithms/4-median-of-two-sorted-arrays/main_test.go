package main

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	if findMedianSortedArrays([]int{1, 2}, []int{3, 4}) != 2.5 {
		t.Failed()
	}
	if findMedianSortedArrays([]int{}, []int{3, 4}) != 3.5 {
		t.Failed()
	}
	if findMedianSortedArrays([]int{1}, []int{}) != 1 {
		t.Failed()
	}
	if findMedianSortedArrays([]int{1, 1, 2, 3, 5, 7, 7, 8}, []int{2, 2, 4, 5, 5, 6, 7, 7}) != 4.5 {
		t.Failed()
	}
}
