package BinarySearch

import (
	"dsa-revision/MinMaxHelpers"
	"math"
)

// link - https://practice.geeksforgeeks.org/problems/k-th-element-of-two-sorted-array1317/1

func KthElementOfSortedArrays(nums1, nums2 []int, k int) int {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	if k > m+n {
		return 0.0
	}
	// low 0 means we can 0 elements from nums1 
	// high m means we can pick all (m elements) from nums1
	low, high := 0, m
	for low <= high {
		// mid1, mid2 states number of elements. not the index until which
		mid1 := (low + high) >> 1
		mid2 := k - mid1
		l1, l2, r1, r2 := math.MinInt64, math.MinInt64, math.MaxInt64, math.MaxInt64
		if mid1 > 0 {
			l1 = nums1[mid1-1]
		}
		if mid1 < m {
			r1 = nums1[mid1]
		}
		if mid2 > 0 {
			l2 = nums2[mid2-1]
		}
		if mid2 < n {
			r2 = nums2[mid2]
		}
		if l1 <= r2 && l2 <= r1 {
			return MinMaxHelpers.Max(l1,l2)
		}
		if l1 > r2 {
			high = mid1-1
		} else {
			low = mid1+1
		}
	}
	return 0.0
}