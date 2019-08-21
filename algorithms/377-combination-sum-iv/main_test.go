package main

import (
	"testing"
)

func TestCombinationSum(t *testing.T) {
	if 6 != combinationSum4([]int{2, 3, 5}, 8) {
		t.Fatal()
	}
	if 7 != combinationSum4([]int{1, 2, 3}, 4) {
		t.Fatal()
	}
	if 39882198 != combinationSum4([]int{4, 2, 1}, 32) {
		t.Fatal()
	}
	if 1132436852 != combinationSum4([]int{2, 1, 3}, 35) {
		t.Fatal()
	}
}
