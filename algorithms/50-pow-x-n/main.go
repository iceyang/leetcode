package main

func MyPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / MyPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	res := MyPow(x, n/2)
	if n%2 == 0 {
		return res * res
	} else {
		return x * res * res
	}
}

// version 3
// func MyPow(x float64, n int) float64 {
// 	if n < 0 {
// 		return 1 / MyPow(x, -n)
// 	}
// 	if n == 0 {
// 		return 1
// 	}
// 	if n == 1 {
// 		return x
// 	}
// 	if n%2 == 0 {
// 		res := MyPow(x, n/2)
// 		return res * res
// 	} else {
// 		return MyPow(x, n/2) * MyPow(x, n/2+n%2)
// 	}
// }

// func MyPow(x float64, n int) float64 {
// 	if n < 0 {
// 		return 1 / MyPow(x, -n)
// 	}
// 	var res float64
// 	res = 1
// 	for i := n; i > 0; i /= 2 {
// 		if i%2 != 0 {
// 			res *= x
// 		}
// 		x *= x
// 	}
// 	return res
// }

// version 2
// func MyPow(x float64, n int) float64 {
// 	if n < 0 {
// 		return 1 / MyPow(x, -n)
// 	}
// 	pow := make([]float64, n+1)
// 	pow[0] = 1
// 	pow[1] = x
// 	i := 2
// 	for i <= n {
// 		pow[i] = pow[i/2] * pow[i/2+i%2]
// 		i++
// 	}
// 	return pow[n]
// }

// version 1
// func MyPow(x float64, n int) float64 {
// 	if n < 0 {
// 		return 1 / MyPow(x, -n)
// 	}
// 	if n == 0 {
// 		return 1
// 	}
// 	return x * MyPow(x, n-1)
// }
