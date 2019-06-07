package main

import (
	"strconv"
)

func getPermutation(n int, k int) string {
	k -= 1
	numbers := make([]int, n)
	for i := range numbers {
		numbers[i] = i + 1
	}
	count := 1
	for i := 2; i <= n; i++ {
		count *= i
	}
	str := ""
	for n > 0 {
		count /= n
		i := k / count
		k %= count
		n--
		str += strconv.Itoa(numbers[i])
		numbers = append(numbers[:i], numbers[i+1:]...)
	}
	return str
}
