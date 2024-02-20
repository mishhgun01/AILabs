package state

import "AILabs/lab1/pkg/consts"

// swap - Перестановка элементов матрицы.
func swap(i, j int, pos string, oldMatrix [3][3]int) [3][3]int {
	matrix := oldMatrix
	switch pos {
	case consts.LEFT:
		if j == 0 {
			return oldMatrix
		}
		matrix[i][j-1], matrix[i][j] = matrix[i][j], matrix[i][j-1]
		return matrix
	case consts.RIGHT:
		if j == 2 {
			return oldMatrix
		}
		matrix[i][j+1], matrix[i][j] = matrix[i][j], matrix[i][j+1]
		return matrix
	case consts.UP:
		if i == 0 {
			return oldMatrix
		}
		matrix[i][j], matrix[i-1][j] = matrix[i-1][j], matrix[i][j]
		return matrix
	case consts.DOWN:
		if i == 2 {
			return oldMatrix
		}
		matrix[i][j], matrix[i+1][j] = matrix[i+1][j], matrix[i][j]
		return matrix
	}

	return matrix
}
