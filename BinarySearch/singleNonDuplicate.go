package BinarySearch

// link - https://takeuforward.org/data-structure/search-single-element-in-a-sorted-array/

func SingleNonDuplicate(arr []int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	if n==1 {
		return arr[0]
	}
	if arr[0] != arr[1] {
		return arr[0]
	}
	if arr[n-2] != arr[n-1] {
		return arr[n-1]
	}
	low, high := 1, n-2
	for low <= high {
		mid  := (low+high) >> 1
		if arr[mid] != arr[mid+1] && arr[mid] != arr[mid-1] {
			return arr[mid]
		}
		// our mid is at left of target. so we cut left of mid 
		if (mid%2==0 && arr[mid]==arr[mid+1]) || (mid%2==1 && arr[mid]==arr[mid-1]) {
			low = mid+1
		} else {
			high = mid-1
		}
	}
	return -1
}