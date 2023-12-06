package RotateMatrix

func Rotate(matrix [][]int) {
	// given nxn matrix
	n := len(matrix)

	// find the transponse of matrix

	for row := 0; row < n; row ++ {
		// check only upper triangle of the matrix, ignore diagonal
		for col := row+1; col < n; col ++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}

	// reverse every row
	for row := 0; row < n; row ++ {
		colStart, colEnd := 0, n-1
		for colStart < colEnd {
			matrix[row][colStart], matrix[row][colEnd] = matrix[row][colEnd], matrix[row][colStart]
			colStart ++ 
			colEnd --
		}
	}
}