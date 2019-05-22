package main

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	num := ClimbStairs(10)
	if num != 89 {
		t.Fail()
	}
}
