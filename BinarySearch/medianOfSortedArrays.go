package BinarySearch

import (
	"dsa-revision/MinMaxHelpers"
	"math"
)

// link - https://leetcode.com/problems/median-of-two-sorted-arrays/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	low, high := 0, m
	for low <= high {
		mid1 := (low + high) >> 1
		mid2 := (m+n+1)>>1 - mid1
		l1 := math.MinInt64
		if mid1 > 0 {
			l1 = nums1[mid1-1]
		}
		r1 := math.MaxInt64
		if mid1 < m {
			r1 = nums1[mid1]
		}
		l2 := math.MinInt64
		if mid2 > 0 {
			l2 = nums2[mid2-1]
		}
		r2 := math.MaxInt64
		if mid2 < n {
			r2 = nums2[mid2]
		}

		if l1 <= r2 && l2 <= r1 {
			if (m+n)%2 == 0 {
				return (float64(MinMaxHelpers.Max(l1, l2)) + float64(MinMaxHelpers.Min(r1, r2))) / 2.0
			}
			return float64(MinMaxHelpers.Max(l1, l2))
		} else if l1 > r2 {
			high = mid1 - 1
		} else {
			low = mid1 + 1
		}
	}
	return 0.0

}
