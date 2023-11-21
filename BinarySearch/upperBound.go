package BinarySearch

// In a given sorted array, find number of elements <= given key
func UpperBoundCount(arr []int, size int, key int) int {
	left, right, mid := 0, size-1, 0

	for left <= right {
		mid = (left + right) >> 1
		if arr[mid] == key {
			// Find the last occurrence of key
			for mid+1 < size && arr[mid+1] == key {
				mid += 1
			}
			break
		} else if arr[mid] > key {
			right = mid-1
		} else {
			left = mid+1
		}
	}
	// If key is not found, mid is the insertion point
    // before key
	for mid > -1 && arr[mid]>key {
		mid -= 1
	}
	return mid+1
}