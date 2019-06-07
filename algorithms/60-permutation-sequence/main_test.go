package main

import (
	"testing"
)

func TestGetPermutation(t *testing.T) {
	if getPermutation(3, 3) != "213" {
		t.Fail()
	}
	if getPermutation(5, 17) != "14523" {
		t.Fail()
	}
	if getPermutation(7, 50) != "1253476" {
		t.Fail()
	}
	if getPermutation(9, 60484) != "261345897" {
		t.Fail()
	}
}
