package BinarySearch

// link - https://leetcode.com/problems/search-in-rotated-sorted-array-ii/

func SearchInRotatedArrayWithRepeatedElements(nums []int, target int) bool {
    n := len(nums)
    low, high := 0, n-1

	// ex - nums = [1,1,1,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1]. low, mid, high elem = 1. We can't decide upon which half to leave.
	// hence in cases where tails are extented equally to heads, we can remove equivalent part of head to remove this confusion 
	// without impacting answer

    for low < n-1 && nums[low]==nums[n-1]{
        low ++
    }
    for low <= high {
        mid := (low+high) >> 1
        if nums[mid] == target {
            return true
        }
        if nums[low] <= nums[mid] {
            if nums[low]<=target && target < nums[mid] {
                high = mid-1
            } else {
                low = mid+1
            }
        } else {
            if nums[mid] < target && target <= nums[high] {
                low = mid+1
            } else {
                high = mid-1
            }
        }
    }
    return false
}