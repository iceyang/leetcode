package main

import "testing"

func TestWordBreak(t *testing.T) {
	wordBreak("leetcode", []string{"leet", "code"})
	wordBreak("applepenapple", []string{"apple", "pen"})
	wordBreak("applepenapple", []string{"applepen"})
	wordBreak("applepenapple", []string{"apple", "pen", "applepen"})
	wordBreak("aaaa", []string{"a", "aa"})
	wordBreak("", []string{""})
	wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"})
}
