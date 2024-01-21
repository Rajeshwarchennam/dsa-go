package SlowAndFastPointers


// https://leetcode.com/problems/find-the-duplicate-number/description/ 


// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

// There is only one repeated number in nums, return this repeated number. 


func findDuplicate(nums []int) int {
    // fast pointer goes twice as fast as slow pointer
	fast, slow := nums[0], nums[0]
	for {
		fast, slow = nums[nums[fast]], nums[slow]
		if fast == slow {
			break
		}
	}

	// start another pointer at nums[0] and go slow 
	// merging of slow and this another pointer gives the duplicate element
	anotherSlowPointer := nums[0]
	for anotherSlowPointer != slow {
		anotherSlowPointer, slow = nums[anotherSlowPointer], nums[slow]
	}
	return slow
} 

