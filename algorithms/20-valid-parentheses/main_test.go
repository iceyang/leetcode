package main

import "testing"

func TestIsValid(t *testing.T) {
	if !isValid("") {
		t.Fail()
	}
	if !isValid("()") {
		t.Fail()
	}
	if isValid("(]") {
		t.Fail()
	}
	if !isValid("()[]{}") {
		t.Fail()
	}
	if !isValid("{[]}") {
		t.Fail()
	}
}
