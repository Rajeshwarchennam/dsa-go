package SlowAndFastPointers

// link - https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/ 

// Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
// Return k.

func removeDuplicates(nums []int) int {
    i := 0
    for j := range nums {
        if nums[i] != nums[j] {
            i += 1
            nums[i] = nums[j]
        }
    }
    return i + 1
}