package SetMatrixZeroes

//link - https://leetcode.com/problems/set-matrix-zeroes/description/

func SetZeroes(matrix [][]int) {
	isFirstRowZero, isFirstColZero := false, false
	rows, columns := len(matrix), len(matrix[0])

	// check if first row is zero
	for j := 0; j < columns; j++ {
		if matrix[0][j] == 0 {
			isFirstRowZero = true
			break
		}
	}

	// check if first col is zero
	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			isFirstColZero = true
			break
		}
	}

	// we check for all rows and cols except first row and first column
	// we will put first row and first column as a placeholder to tell if that
	// whole row or column will be replaced by zeroes
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}

	// if either matrix[0][j] or matrix[i][0] is zero, we can state matrix[i][j] is zero
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// if firstrow or firstcolumn needs to be zero valued, we change them

	if isFirstRowZero {
		for j := 0; j < columns; j++ {
			matrix[0][j] = 0
		}
	}

	if isFirstColZero {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
}
