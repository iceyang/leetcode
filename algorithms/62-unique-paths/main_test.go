package main

import "testing"

func TestUniquePaths(t *testing.T) {
	if uniquePaths(0, 0) != 0 {
		t.Fail()
	}
	if uniquePaths(1, 1) != 1 {
		t.Fail()
	}
	if uniquePaths(3, 2) != 3 {
		t.Fail()
	}
	if uniquePaths(7, 3) != 28 {
		t.Fail()
	}
}
