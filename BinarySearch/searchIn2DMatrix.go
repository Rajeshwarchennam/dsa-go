package BinarySearch

// link -https://leetcode.com/problems/search-a-2d-matrix/

func SearchMatrix(matrix [][]int, target int) bool {
    rows, columns := len(matrix), len(matrix[0])
    low, high := 0, rows*columns-1
    for low <= high {
        mid := (low+high) >>  1
        midElemRow, midElemCol := mid/columns, mid%columns
        midElem := matrix[midElemRow][midElemCol]
        if midElem == target {
            return true
        }
        if midElem > target {
            high = mid-1
        } else {
            low = mid+1
        }
    }
    return false
}