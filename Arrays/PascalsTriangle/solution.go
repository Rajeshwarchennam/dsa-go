package PascalsTriangle

// link - https://leetcode.com/problems/pascals-triangle/description/

func Generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	pascalTriangle := make([][]int, 0, numRows)
	pascalTriangle = append(pascalTriangle, []int{1})
	if numRows == 1 {
		return pascalTriangle
	}
	for i := 2; i <= numRows; i++ {
		currentRow := make([]int, 0, i)
		// pad current row with 1 at start
		currentRow = append(currentRow, 1)
		// sum of last row in pascalTriangle consecutive elements
		for j := 0; j < i-2; j++ {
			currentRow = append(currentRow, pascalTriangle[i-2][j]+pascalTriangle[i-2][j+1])
		}
		// pad current row with 1 at end
		currentRow = append(currentRow, 1)
		pascalTriangle = append(pascalTriangle, currentRow)
	}
	return pascalTriangle
}