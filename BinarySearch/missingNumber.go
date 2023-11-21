package BinarySearch

import "sort"

// link - https://leetcode.com/problems/missing-number/

func MissingNumber(nums []int) int {
	sort.Ints(nums)
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		if mid == nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
