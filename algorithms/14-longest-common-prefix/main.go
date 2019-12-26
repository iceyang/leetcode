package main

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	minLength := len(strs[0])
	for i := 1; i < length; i++ {
		if minLength > len(strs[i]) {
			minLength = len(strs[i])
		}
	}
	cur := 0
	for true {
		if cur == minLength {
			break
		}
		char := strs[0][cur]
		for j := 1; j < length; j++ {
			if char != strs[j][cur] {
				return strs[0][:cur]
			}
		}
		cur++
	}
	return strs[0][:cur]
}
