package main

var letters = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

var result []string

func bt(str string, digits string) {
	if len(digits) == 0 {
		result = append(result, str)
		return
	}
	number := digits[0]
	for _, letter := range letters[number] {
		bt(str+string(letter), digits[1:])
	}
}

func letterCombinations(digits string) []string {
	result = []string{}
	if len(digits) == 0 {
		return result
	}
	bt("", digits)
	return result
}
