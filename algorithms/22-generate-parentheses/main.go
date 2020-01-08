package main

var result [][]string

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	result = make([][]string, n+1)
	result[0] = []string{""}
	result[1] = []string{"()"}
	for i := 2; i <= n; i++ {
		res := make([]string, 0, 2*i)
		for j := 0; j < i; j++ {
			strs1 := result[j]
			strs2 := result[i-1-j]
			for _, str1 := range strs1 {
				for _, str2 := range strs2 {
					res = append(res, "("+str1+")"+str2)
				}
			}
		}
		result[i] = res
	}
	return result[n]
}
