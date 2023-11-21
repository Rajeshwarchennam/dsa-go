package BinarySearch

import "math"

// link - https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/

func MinumumInRotatedArray(nums []int) int {
	low, high := 0, len(nums)-1
	ans := math.MaxInt64
	for low <= high {
		mid := (low+high) >> 1
		if nums[low] <= nums[mid] {
			// if left part is sorted, we can check in right part if answer present. But before leaving left part we need to track min in left part i.e nums[low]
			// example- [1,2,3,4,5] here [1,2,3] are sorted, so we can check ans in second part but make sure we track min in left as it can be answer
			ans = min(ans, nums[low])
			low = mid+1
		} else {
			// example - [5,4,1,2,3] here if we take right part [1,2,3]. before skipping this sorted part, we need to track min of right i.e nums[mid] coz that might be 
			// possible answer
			ans = min(ans, nums[mid])
			high = mid-1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {return a}
	return b 
}