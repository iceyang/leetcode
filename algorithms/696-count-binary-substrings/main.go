package main

func countBinarySubstrings(s string) int {
	counts := []int{}
	length := len(s)
	if length == 0 {
		return 1
	}
	for i := 0; i < length; i++ {
		c := s[i]
		count := 1
		for i < length-1 && s[i+1] == c {
			count += 1
			i++
		}
		counts = append(counts, count)
	}
	result := 0
	for i := len(counts) - 1; i > 0; i-- {
		if counts[i] > counts[i-1] {
			result += counts[i-1]
		} else {
			result += counts[i]
		}
	}
	return result
}
