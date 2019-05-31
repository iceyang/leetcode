package main

import (
	"testing"
)

func TestMaxSatisfied(t *testing.T) {
	if 16 != maxSatisfied(
		[]int{1, 0, 1, 2, 1, 1, 7, 5},
		[]int{0, 1, 0, 1, 0, 1, 0, 1},
		3,
	) {
		t.Fail()
	}
	if 0 != maxSatisfied(
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
		0,
	) {
		t.Fail()
	}
	if 0 != maxSatisfied(
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
		[]int{1, 1, 1, 1, 1, 1, 1, 1},
		0,
	) {
		t.Fail()
	}
	if 23 != maxSatisfied(
		[]int{4, 7, 0, 11, 5, 7, 0, 10},
		[]int{1, 1, 1, 1, 1, 1, 1, 1},
		4,
	) {
		t.Fail()
	}
	if 33 != maxSatisfied(
		[]int{4, 7, 0, 11, 5, 7, 0, 10},
		[]int{1, 1, 1, 1, 1, 1, 1, 1},
		5,
	) {
		t.Fail()
	}
	if 44 != maxSatisfied(
		[]int{4, 7, 0, 11, 5, 7, 0, 10},
		[]int{1, 1, 1, 1, 1, 1, 1, 1},
		8,
	) {
		t.Fail()
	}
	if 44 != maxSatisfied(
		[]int{4, 7, 0, 11, 5, 7, 0, 10},
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
		0,
	) {
		t.Fail()
	}
	if 33 != maxSatisfied(
		[]int{4, 7, 0, 11, 5, 7, 0, 10},
		[]int{1, 1, 0, 0, 0, 1, 0, 1},
		5,
	) {
		t.Fail()
	}
	if 27 != maxSatisfied(
		[]int{4, 7, 0, 11, 5, 7, 0, 10},
		[]int{1, 1, 0, 0, 0, 1, 0, 1},
		2,
	) {
		t.Fail()
	}
}
