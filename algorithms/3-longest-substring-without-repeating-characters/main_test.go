package main

import (
	"testing"
)

func assert(expect int, actual int, t *testing.T) {
	if expect == actual {
		return
	}
	t.Errorf("expected %d but got %d\n", expect, actual)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	assert(3, lengthOfLongestSubstring("abcabcbb"), t)
	assert(1, lengthOfLongestSubstring("bbbbb"), t)
	assert(3, lengthOfLongestSubstring("pwwkew"), t)
	assert(7, lengthOfLongestSubstring("abcdefga"), t)
	assert(6, lengthOfLongestSubstring("abcdafga"), t)
	assert(5, lengthOfLongestSubstring("abcdaf"), t)
	assert(5, lengthOfLongestSubstring("abcdaffffaj"), t)
	assert(5, lengthOfLongestSubstring("abcdaffffajlkje"), t)
}
