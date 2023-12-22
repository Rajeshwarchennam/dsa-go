package Recursion

import "sort"

// link - https://practice.geeksforgeeks.org/problems/subset-sums2234/1


// given a list of nums, return sums of all subets in it
func SubsetSums(nums []int, size int) []int {
	res := make([]int, 0, size)
	var recFun func(currIndex int, currSum int)
	recFun = func(currIndex int, currSum int) {
		if currIndex == size {
			res = append(res, currSum)
			return
		}
		recFun(currIndex+1, currSum+nums[currIndex])
		recFun(currIndex+1, currSum)
	}
	recFun(0, 0)
	sort.Ints(res)
	return res
}