package Recursion

// link - https://leetcode.com/problems/n-queens/description
import "strings"

func SolveNQueens(n int) [][]string {
	res := make([][]string, 0, 5)
	transitionList := make([]int, 0, n)
	var recFun func(int, []int)
	recFun = func(currRow int, possilbeColumns []int) {
		// only if it reaches n after filling till n-1, we can add this to our result
		if currRow == n {
			temp := make([]string, 0, len(transitionList))
			// for each val in transitionList we are making a string like "..Q." (val=2, n=4)
			for _, val := range transitionList {
				currRowString := strings.Repeat(".", val) + "Q" + strings.Repeat(".", n-val-1)
				temp = append(temp, currRowString)
			}
			res = append(res, temp)
			return
		}
		// possible columns consists of columns that are allowed by checking not previously used columns
		// this doesn't check the diagnol condition. Hence isSafe check is needed
		for i := range possilbeColumns {
			// for a given (not used column till now in other rows), we can check if this [row][column] is not diagonal to any queen
			// if none of them is diagonally possible, we move back to parent, without proceeding further. Hence discarding our current route.
			if !isSafeToPlaceQueen(transitionList, currRow, possilbeColumns[i], n) {
				continue
			}
			transitionList = append(transitionList, possilbeColumns[i])
			// we remove the ithElement from possibleColumns as we already used that.
			recFun(currRow+1, removeithElementFromList(possilbeColumns, i))
			transitionList = transitionList[:len(transitionList)-1]
		}
	}
	allColumns := make([]int, n)
	for i := 0; i < n; i++ {
		allColumns[i] = i
	}
	recFun(0, allColumns)
	return res
}

func isSafeToPlaceQueen(transitionList []int, currentRow, currentCol int, n int) bool {
	i, j := currentRow, currentCol
	// checking left upper diagonal
	for i > 0 && j > 0 {
		i -= 1
		j -= 1
		if transitionList[i] == j {
			return false
		}
	}
	// checking right upper diagonal
	i, j = currentRow, currentCol
	for i > 0 && j < n-1 {
		i -= 1
		j += 1
		if transitionList[i] == j {
			return false
		}
	}
	return true
}