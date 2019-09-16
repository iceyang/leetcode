package main

var NUMS = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var ROMANS = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func intToRoman(num int) string {
	result := ""
	for i := 0; i < len(ROMANS); i++ {
		for num >= NUMS[i] {
			result += ROMANS[i]
			num -= NUMS[i]
		}
	}
	return result
}
