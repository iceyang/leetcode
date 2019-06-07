package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if !isPalindrome(121) {
		t.Fail()
	}
	if isPalindrome(-121) {
		t.Fail()
	}
	if isPalindrome(10) {
		t.Fail()
	}
	if !isPalindrome(123321) {
		t.Fail()
	}
}
