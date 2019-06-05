package main

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	length := 0
	startIndex := 0
	maxLength := 0
	for index, char := range s {
		i, exist := charMap[char]
		if exist && i+1 > startIndex {
			startIndex = i + 1
		}
		length = index - startIndex + 1
		charMap[char] = index
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
