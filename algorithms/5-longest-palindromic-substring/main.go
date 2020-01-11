package main

func isPalindrome(s *string) bool {
	for l, r := 0, len(*s)-1; l < r; {
		if (*s)[l] != (*s)[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func longestPalindrome(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}
	pre := longestPalindrome(s[:length-1])
	for i, max := 0, length-len(pre); i < max; i++ {
		str := s[i:]
		if isPalindrome(&str) {
			return str
		}
	}
	return pre
}
