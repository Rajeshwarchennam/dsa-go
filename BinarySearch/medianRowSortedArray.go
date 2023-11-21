package BinarySearch

import "dsa-revision/MinMaxHelpers"

// link - https://takeuforward.org/data-structure/median-of-row-wise-sorted-matrix/

func MedianOfRowWiseSortedMatrix(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return -1
	}
	m := len(matrix[0])
	if m == 0 {
		return -1
	}
	low, high := matrix[0][0], matrix[0][m-1]
	for i := 1; i < n; i++ {
		low = MinMaxHelpers.Min(low, matrix[i][0])
		high = MinMaxHelpers.Max(high, matrix[i][m-1])
	}
	medianPos := (n*m)/2 + 1
	for low <= high {
		mid := (low+high) >> 1
		count := 0
		for i := 0; i<n; i++ {
			count += UpperBoundCount(matrix[i], m, mid)
		}
		// need to check the least element that satisfies <= medianPosition condition 
		// [1,2,4,4,7,9,10] because 5 also satisifies condition but we need to move to the point which is present in matrix using fact that beyond that point condition fails. 
		if count == medianPos {
			high = mid-1
		} else if count > medianPos {
			high = mid-1
		} else {
			low = mid+1
		}
	}
	return low
}