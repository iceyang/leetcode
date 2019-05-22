package main

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	num := ClimbStairs(10)
	if num != 89 {
		t.Fail()
	}
}
