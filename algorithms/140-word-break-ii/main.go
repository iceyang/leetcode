package main

var result []string

func back(words [][]string, cur int, str string) {
	for _, word := range words[cur] {
		lenOfWord := len(word)
		if cur-lenOfWord == 0 {
			finalWord := word + " " + str
			result = append(result, finalWord[:len(finalWord)-1])
			continue
		}
		back(words, cur-lenOfWord, word+" "+str)
	}
}

func dp(s string, wordDict []string) []string {
	length := len(s)
	res := make([][]string, length+1)
	for i := 1; i <= length; i++ {
		for _, word := range wordDict {
			lenOfWord := len(word)
			if lenOfWord > i {
				continue
			}
			startIndex := i - lenOfWord
			if word == s[startIndex:i] {
				if lenOfWord == i {
					res[i] = append(res[i], word)
					continue
				}
				if len(res[startIndex]) == 0 {
					continue
				}
				res[i] = append(res[i], word)
			}
		}
	}
	if len(res[length]) == 0 {
		return []string{}
	}
	back(res, length, "")
	return result
}

func wordBreak(s string, wordDict []string) []string {
	if len(s) == 0 {
		return []string{""}
	}
	if len(wordDict) == 0 {
		return []string{}
	}
	return dp(s, wordDict)
}
