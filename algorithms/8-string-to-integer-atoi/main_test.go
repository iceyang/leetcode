package main

import (
	"testing"
)

func TestMyAtoi(t *testing.T) {
	if myAtoi("42") != 42 {
		t.Fail()
	}
	if myAtoi("4193 with words") != 4193 {
		t.Fail()
	}
	if myAtoi("words and 987") != 0 {
		t.Fail()
	}
	if myAtoi("2147483648") != 2147483647 {
		t.Fail()
	}
	if myAtoi("-2147483649") != -2147483648 {
		t.Fail()
	}
	if myAtoi("0-1") != 0 {
		t.Fail()
	}
}
