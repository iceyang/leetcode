package main

var MAX_VALUE = 2147483647
var MIN_VALUE = -2147483648

func reverse(x int) int {
	result := 0
	for x != 0 {
		remain := x % 10
		x /= 10
		if result > MAX_VALUE/10 || (result == MAX_VALUE/10 && remain > 7) ||
			result < MIN_VALUE/10 || (result == MIN_VALUE/10 && remain < -8) {
			return 0
		}
		result = result*10 + remain
	}
	return result
}
