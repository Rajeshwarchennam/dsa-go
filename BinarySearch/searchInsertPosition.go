package BinarySearch 

// link - https://leetcode.com/problems/search-insert-position/description/

func SearchInsertPosition(nums []int, target int) int {
    low, high := 0, len(nums)-1
    for low <= high {
        mid := (low+high) >> 1
        if nums[mid] == target {
            return mid
        }
        if nums[mid] > target {
            high = mid-1
        } else {
            low = mid+1
        }
    }
    return low
}