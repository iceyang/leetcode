package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	x2 := x
	y := 0
	for x2 > 0 {
		y = y*10 + x2%10
		x2 /= 10
	}
	return x == y
}
