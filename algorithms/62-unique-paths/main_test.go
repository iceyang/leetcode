package main

import "testing"

func TestUniquePaths(t *testing.T) {
	if uniquePaths(7, 3) != 28 {
		t.Fail()
	}
	if uniquePaths(0, 0) != 0 {
		t.Fail()
	}
	if uniquePaths(1, 1) != 1 {
		t.Fail()
	}
	if uniquePaths(3, 2) != 3 {
		t.Fail()
	}
	if uniquePaths(20, 20) != 35345263800 {
		t.Fail()
	}
	if uniquePaths(30, 23) != 156077261327400 {
		t.Fail()
	}
}
