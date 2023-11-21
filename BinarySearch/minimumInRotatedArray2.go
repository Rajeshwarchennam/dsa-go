package BinarySearch

import "math"

// link - https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/description/

func MinimumInRotatedArrayWithDuplicates(nums []int) int {
	n := len(nums)
	low, high := 0, n-1

	// ex - nums = [1,1,1,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1]. low, mid, high elem = 1. We can't decide upon which half to leave.
	// hence in cases where tails are extented equally to heads, we can remove equivalent part of head to remove this confusion 
	// without impacting answer
	for low < n-1 && nums[low] == nums[n-1] {
		low += 1
	}
	ans := math.MaxInt64
	for low <= high {
		mid := (low + high) >> 1
		if nums[low] <= nums[mid] {
			ans = min(ans, nums[low])
			low = mid + 1
		} else {
			ans = min(ans, nums[mid])
			high = mid - 1
		}
	}
	return ans
}
