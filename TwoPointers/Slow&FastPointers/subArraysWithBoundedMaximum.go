package SlowAndFastPointers

// link - https://leetcode.com/problems/number-of-subarrays-with-bounded-maximum/description/ 

// Given an integer array nums and two integers left and right, return the number of contiguous non-empty subarrays such that the value of the maximum array element in that subarray is in the range [left, right].


func numSubarrayBoundedMax(nums []int, left int, right int) int {
    // two pointers usage
    start, end, ans  := 0, 0, 0
    subArraysThatCanEndWithEnd := 0
    for end < len(nums) {
        // if nums[end] lies in [left, right], nums[end] can be included in all the subarrays 
        // that can end with "end" index and can start all the way from "start" index to 
        // "end" index (end index itself can start a new subarray)
        if nums[end] >= left && nums[end] <= right {
            subArraysThatCanEndWithEnd = end-start+1
            ans += subArraysThatCanEndWithEnd
        } else if nums[end] < left {
            // if nums[end] is less than left, we can include in all preceeding subarrays but 
            // we cannot start a subarray with this element (atleast without knowing the next element)
            ans += subArraysThatCanEndWithEnd
        } else {
            // we need to reset our start, when we see an eleemnt greater than right
            start = end+1
            subArraysThatCanEndWithEnd = 0 
        }
        end ++
    }
    return ans
}
