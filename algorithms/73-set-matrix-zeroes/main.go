package main

func setZeroes(matrix [][]int) {
	rowLen := len(matrix)
	if rowLen <= 0 {
		return
	}
	columnLen := len(matrix[0])
	if columnLen <= 0 {
		return
	}
	rows := map[int]bool{}
	columns := map[int]bool{}
	for i := rowLen - 1; i >= 0; i-- {
		for j := columnLen - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				rows[i] = true
				columns[j] = true
			}
		}
	}
	for row, _ := range rows {
		for j := columnLen - 1; j >= 0; j-- {
			matrix[row][j] = 0
		}
	}
	for column, _ := range columns {
		for i := rowLen - 1; i >= 0; i-- {
			matrix[i][column] = 0
		}
	}
}
