package main

var dict = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func romanToInt(s string) int {
	res := 0
	for i, length := 0, len(s); i < length; {
		if i+1 < length {
			v, ok := dict[s[i:i+2]]
			if ok {
				res += v
				i += 2
				continue
			}
		}
		v, _ := dict[s[i:i+1]]
		res += v
		i += 1
	}
	return res
}
