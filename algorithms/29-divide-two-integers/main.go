package main

// func divide(dividend int, divisor int) int {
// 	positive := true
// 	if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
// 		positive = false
// 	}
// 	if dividend > 0 {
// 		dividend = -dividend
// 	}
// 	if divisor > 0 {
// 		divisor = -divisor
// 	}
// 	result := 0
// 	for true {
// 		dividend -= divisor
// 		if dividend > 0 {
// 			break
// 		}
// 		result += 1
// 	}
// 	if !positive {
// 		return -result
// 	}
// 	return result
// }

func div(dividend, divisor int) int {
	if dividend < divisor {
		return 0
	}
	result := 1
	divisorBak := divisor
	for divisor+divisor <= dividend {
		result += result
		divisor += divisor
	}
	return result + div(dividend-divisor, divisorBak)
}

func divide(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == -2147483648 {
			return 2147483647
		}
		return -dividend
	}
	positive := true
	if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
		positive = false
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	if !positive {
		return -div(dividend, divisor)
	}
	return div(dividend, divisor)
}
