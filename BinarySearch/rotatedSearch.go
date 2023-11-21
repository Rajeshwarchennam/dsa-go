package BinarySearch

// link - https://leetcode.com/problems/search-in-rotated-sorted-array/description/

func SearchInRotatedArray(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low+high) >> 1
		if arr[mid] == target {
			return mid
		}
		if arr[low] <= arr[mid] {
			if arr[low] <= target && target < arr[mid] {
				high = mid-1
			} else {
				low = mid+1
			}
		} else {
			if arr[mid] < target && target <= arr[high] {
				low = mid+1
			} else {
				high = mid-1
			}
		}
	}
	return -1
}