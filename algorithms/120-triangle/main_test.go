package main

import "testing"

func TestMinimumTotal(t *testing.T) {
	triangle := make([][]int, 4)
	triangle[0] = []int{2}
	triangle[1] = []int{3, 4}
	triangle[2] = []int{6, 5, 7}
	triangle[3] = []int{4, 1, 8, 3}
	if minimumTotal(triangle) != 11 {
		t.Fail()
	}
}
