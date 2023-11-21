package Recursion

// link - https://leetcode.com/problems/sudoku-solver/description/

// example board - 
// board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
// 		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
// 		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
// 		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
// 		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
// 		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
// 		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
// 		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
// 		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

func SolveSudoku(board [][]byte) {
	var recFun func() bool
	recFun = func() bool {
		// check for the occurence of slot
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					// check which element fits from 1 to 9
					for elem := byte('1'); elem <= '9'; elem++ {
						// for each element, check if it is valid in that place
						if isValid(board, i, j, elem) {
							// if valid put that element in that slot
							board[i][j] = elem
							// call the recursive function again. Which states if by putting that element we are getting a valid sudoku
							if recFun() {
								// if it's valid we can return true. This would propagate to the parent. We end up to that solution.
								// as in question it states only one valid solution is possible.
								return true
							} else {
								// if by putting that element, we not able to solve sudoku. put back that element to "."
								board[i][j] = '.'
							}
						}
					}
					// though after checking all the possible values and we are not able to come with valid sudoku. We can tell that 
					// with current state of board we won't be able to solve. So Parent would clear that lastly added elem and try with 
					// new element
					return false
				}
			}
		}
		// if we don't find any empty slot. Our board is solved. 
		return true
	}
	recFun()
}

func isValid(board [][]byte, row, column int, elem byte) bool {
	for i := 0; i < 9; i++ {
		// checking if elem is present in any row of same column
		if board[i][column] == elem {
			return false
		}
		//  any column of same row
		if board[row][i] == elem {
			return false
		}
		// checking in the block.
		// 3*row/3 brings to row batch start
		// 3*column/3 brings to column batch start
		// i/3 lets us find which row in that batch. whereas i%3 computes column count in column batch.
		if board[3*(row/3)+i/3][3*(column/3)+i%3] == elem {
			return false
		}
	}
	return true
}