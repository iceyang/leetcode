package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	haystackLen := len(haystack)
	needleLen := len(needle)
	end := haystackLen - needleLen
	for i := 0; i <= end; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}
