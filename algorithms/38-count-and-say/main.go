package main

func deal(str string) string {
	var (
		cur   byte
		count byte
	)
	count = '1'
	cur = '0'
	res := ""
	for i, l := 0, len(str); i < l; i++ {
		num := str[i]
		if num == cur {
			count++
		} else {
			res += string(count) + string(cur)
			cur = num
			count = '1'
		}
	}
	res += string(count) + string(cur)
	return res[2:]
}

func countAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		str := deal(res)
		res = str
	}
	return res
}
