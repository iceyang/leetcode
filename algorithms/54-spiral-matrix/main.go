package main

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	if n == 0 {
		return []int{}
	}
	record := make([][]bool, m)
	for index, _ := range record {
		record[index] = make([]bool, n)
	}

	r := 0
	c := 0
	// 决定下一步的方向
	di := 0
	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}

	length := m * n
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = matrix[r][c]
		record[r][c] = true
		cr, cc := r+dr[di], c+dc[di]
		if 0 <= cr && cr < m && 0 <= cc && cc < n && !record[cr][cc] {
			r, c = cr, cc
		} else {
			di = (di + 1) % 4
			r += dr[di]
			c += dc[di]
		}
	}
	return res
}
