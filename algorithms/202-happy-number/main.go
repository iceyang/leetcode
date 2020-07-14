package main

var result map[int]int

func loop(n int) bool {
	if n <= 0 || result[n] != 0 {
		return false
	}
	if n == 1 {
		return true
	}
	result[n] = 1
	n2 := 0
	for n != 0 {
		num := n % 10
		n2 += num * num
		n /= 10
	}
	return loop(n2)
}

func isHappy(n int) bool {
	result = map[int]int{}
	return loop(n)
}
