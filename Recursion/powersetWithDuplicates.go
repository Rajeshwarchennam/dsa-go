package Recursion

import "sort"

// link - https://leetcode.com/problems/subsets-ii/


// given an integer array nums (may have duplicates), return powerset
func SubsetsWithDuplicates(nums []int, size int) [][]int {
	res := make([][]int, 0, size)
    transitionlist := make([]int, 0)
	var recFunc func(index int)
	recFunc = func(index int) {
		//at every recursion call we append current state of transitionList to result 
        temp := make([]int, len(transitionlist))
        copy(temp, transitionlist)
		res = append(res, temp)
		for i := index; i < size; i++ {
			// we add the element at given index (as we see that state first time for a given instance of transitionlist formed) and avoid duplicates from then
			if i != index && nums[i] == nums[i-1] {
				continue
			}
            transitionlist = append(transitionlist, nums[i])
			// we call recFunc right after the next element that we have added. because [1,2] & [1,2,2] are distinct subsets
			recFunc(i+1)
            transitionlist = transitionlist[:len(transitionlist)-1]
		}
	}
	// we need to sort 
	sort.Ints(nums)
	recFunc(0)
	return res
}