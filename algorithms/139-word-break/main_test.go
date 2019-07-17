package main

import "testing"

func TestWordBreak(t *testing.T) {
	if !wordBreak("leetcode", []string{"leet", "code"}) {
		t.Fatal()
	}
	if !wordBreak("applepenapple", []string{"apple", "pen"}) {
		t.Fatal()
	}
	if wordBreak("applepenapple", []string{"applepen"}) {
		t.Fatal()
	}
	if wordBreak("applepenapple", []string{""}) {
		t.Fatal()
	}
	if !wordBreak("", []string{""}) {
		t.Fatal()
	}
	if !wordBreak("", []string{"a"}) {
		t.Fatal()
	}
	if wordBreak("a", []string{}) {
		t.Fatal()
	}
	if wordBreak("a", []string{"ab"}) {
		t.Fatal()
	}
	if wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}) {
		t.Fatal()
	}
}
