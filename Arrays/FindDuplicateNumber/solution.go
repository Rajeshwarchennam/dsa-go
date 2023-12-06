package FindDuplicateNumber 

// link - https://leetcode.com/problems/find-the-duplicate-number/

// nums has n+1 integers. It contains all [1, n] elements and 
// one element is repeated. Find that

func FindDuplicate(nums []int) int {

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