package BinarySearch

// link - https://takeuforward.org/data-structure/nth-root-of-a-number-using-binary-search/

func FindNthRoot(num, n int) int {
	low, high := 1, num
	for low <= high {
		mid := (low+high) >> 1
		status := getStatus(mid, num, n)
		if status == 0 {
			return mid
		}
		if status == 1 {
			high = mid-1
		} else {
			low = mid+1
		}
	}
	return -1
}

func getStatus(curr, num, n int) int {
	ans := 1
	for i:=1 ; i<=n; i++ {
		ans = ans*curr
		if ans > num {
			// 1 means curr is  greater than our answer
			return  1
		}
	}
	if ans == num  {
		// 0 means curr is our answer
		return 0
	}
	// -1 means curr is lesser than our answer
	return  -1
}