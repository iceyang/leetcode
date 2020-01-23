package main

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]byte, len(s))
	length := 0
	for i, l := 0, len(s); i < l; i++ {
		c := s[i]
		if length == 0 {
			stack[length] = c
			length += 1
			continue
		}
		pre := stack[length-1]
		if (c == ')' && pre == '(') || (c == ']' && pre == '[') || (c == '}' && pre == '{') {
			length -= 1
			continue
		}
		stack[length] = c
		length += 1
	}
	return length == 0
}
