package main

func dp(s string, wordDict []string) bool {
	length := len(s)
	res := make([]bool, length+1)
	res[0] = true
	for i := 1; i <= length; i++ {
		for _, word := range wordDict {
			lenOfWord := len(word)
			if lenOfWord > i {
				continue
			}
			res[i] = res[i-lenOfWord] && word == s[i-lenOfWord:i]
			if res[i] {
				break
			}
		}
	}
	return res[length]
}

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	if len(wordDict) == 0 {
		return false
	}
	return dp(s, wordDict)
}
