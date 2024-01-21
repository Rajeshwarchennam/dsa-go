package SlowAndFastPointers

// link - https://leetcode.com/problems/rotating-the-box/description/ 

// You are given an m x n matrix of characters box representing a side-view of a box. Each cell of the box is one of the following:

// A stone '#'
// A stationary obstacle '*'
// Empty '.'
// The box is rotated 90 degrees clockwise, causing some of the stones to fall due to gravity. Each stone falls down until it lands on an obstacle, another stone, or the bottom of the box. Gravity does not affect the obstacles' positions, and the inertia from the box's rotation does not affect the stones' horizontal positions.



func rotateTheBox(input [][]byte) [][]byte {
    m, n := len(input), len(input[0])
	result := make([][]byte, n)
	for i := 0; i < n; i++ {
		result[i] = make([]byte, m)
	}

	for row := 0; row < m; row ++ {
		rotateCurrentRowAndFillInResultBox(row, m, n, input, result)
	}
    return result
}

func rotateCurrentRowAndFillInResultBox(row, rowsCount, columnsCount int,input [][]byte,  result [][]byte) {
	// start here is meant for result filling
    // end here is meant for input evaluation
    start, end := columnsCount-1, columnsCount-1

	resultColumnToBeFilled := rowsCount-row-1

	for end >= 0 {
		if input[row][end] == '#' {
			result[start][resultColumnToBeFilled] = '#'
			start --
		} else if input[row][end] == '*' {
			for start > end {
				result[start][resultColumnToBeFilled] = '.'
				start --
			}
			result[start][resultColumnToBeFilled] = '*'
			start -- 
		}
		end --
	}
	for start > end {
		result[start][resultColumnToBeFilled] = '.'
		start --
	}
}