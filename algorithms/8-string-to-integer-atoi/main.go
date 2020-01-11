package main

import (
	"strings"
)

const INT_MAX = 2147483647
const INT_MIN = -2147483648

func myAtoi(str string) int {
	var positive rune
	var result rune
	str2 := strings.TrimSpace(str)
	for _, char := range str2 {
		if char == '-' || char == '+' {
			if positive != 0 {
				break
			}
			positive = char
			continue
		}
		if char >= 48 && char <= 57 {
			if positive == 0 {
				positive = '+'
			}
			if positive != '-' && result > INT_MAX/10 {
				return INT_MAX
			}
			if positive == '-' && -result < INT_MIN/10 {
				return INT_MIN
			}
			result = result*10 + (char - 48)
			continue
		}
		break
	}
	if positive == '-' {
		if result < 0 {
			return INT_MIN
		}
		return int(-result)
	}
	if result < 0 {
		return INT_MAX
	}
	return int(result)
}
