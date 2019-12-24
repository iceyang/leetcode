package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	longestPalindrome("abba")
	longestPalindrome("aba")
	longestPalindrome("ab")
	longestPalindrome("abcbbbcba")
}
