package main

import "fmt"

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func longestPalindrome(s string) string {
	t := isPalindrome(s)
	fmt.Println(t)
	return s
}
