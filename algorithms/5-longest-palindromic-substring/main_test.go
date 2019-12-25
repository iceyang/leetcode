package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	longestPalindrome("abba")
	longestPalindrome("")
	longestPalindrome("a")
	longestPalindrome("abdcba")
}
