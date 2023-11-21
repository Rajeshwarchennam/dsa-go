package BinarySearch

import "dsa-revision/MinMaxHelpers"

func FindMinimumOfMaximumAllocatedPages(books []int, n, studentsAllowed int) int {
	if studentsAllowed > n {
		return -1
	}
	low := MinMaxHelpers.MinInList(books)
	high := 0 
	for _, pages := range books {
		high += pages
	}
	res := 0
	for low <= high {
		mid := (low+high) >> 1
		students := countStudents(books, n, mid)
		if students > studentsAllowed {
			low = mid+1
		} else {
			res = mid
			high = mid-1
		}
	}
	return res
}

func countStudents(books []int, n, maxAllowedPages int) int {
	currStudent, currStudentPage := 1, 0
	for _, page := range books {
		if currStudentPage + page > maxAllowedPages {
			currStudent += 1
			currStudentPage = page
		} else {
			currStudentPage += page
		}
	}
	return currStudent
}