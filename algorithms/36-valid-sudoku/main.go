package main

func isValidSudoku(board [][]byte) bool {
	rows := [9][10]bool{}
	columns := [9][10]bool{}
	girds := [9][10]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := int(board[i][j]) - 48
			gird := i/3*3 + j/3
			if num == -2 {
				continue
			}
			if rows[i][num] {
				return false
			}
			rows[i][num] = true

			if columns[j][num] {
				return false
			}
			columns[j][num] = true

			if girds[gird][num] {
				return false
			}
			girds[gird][num] = true
		}
	}
	return true
}

// Version One
// func validAndAdd(m *map[byte]bool, num byte) bool {
// 	_, rowExist := (*m)[num]
// 	if rowExist {
// 		return false
// 	}
// 	(*m)[num] = true
// 	return true
// }
//
// func isValidSudoku(board [][]byte) bool {
// 	rows := [9]map[byte]bool{}
// 	columns := [9]map[byte]bool{}
// 	girds := [9]map[byte]bool{}
// 	for i := 0; i < 9; i++ {
// 		rows[i] = map[byte]bool{}
// 		columns[i] = map[byte]bool{}
// 		girds[i] = map[byte]bool{}
// 	}
//
// 	for row := 0; row < 9; row++ {
// 		for column := 0; column < 9; column++ {
// 			num := board[row][column]
// 			gird := row/3*3 + column/3
// 			if num == 46 {
// 				continue
// 			}
// 			if !validAndAdd(&rows[row], num) {
// 				return false
// 			}
// 			if !validAndAdd(&columns[column], num) {
// 				return false
// 			}
// 			if !validAndAdd(&girds[gird], num) {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
