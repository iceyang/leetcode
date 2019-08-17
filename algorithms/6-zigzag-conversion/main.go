package main

func convert(s string, numRows int) string {
	if s == "" {
		return ""
	}
	length := len(s)
	if numRows <= 1 || numRows > length {
		return s
	}
	res := make([][]byte, numRows)
	i := 0
	row := 0
	column := 0
	for i < length {
		res[row] = append(res[row], s[i])
		if column%(numRows-1) != 0 {
			column += 1
			row -= 1
		} else {
			row += 1
		}
		if row == numRows {
			column += 1
			row -= 2
		}
		i++
	}
	result := ""
	for _, r := range res {
		for _, letter := range r {
			if letter == ' ' {
				continue
			}
			result += string(letter)
		}
	}
	return result
}
