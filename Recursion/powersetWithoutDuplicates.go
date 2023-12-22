package Recursion
// link - https://leetcode.com/problems/subsets/

// Given an integer array of unique elements, return all possible substets(power set)

func SubsetsWithoutDuplicates(nums []int, size int) [][]int {
	result := make([][]int, 0, size)
	transitionList := make([]int, 0)
	var recFun func(index int)
	recFun = func(index int) {

		if index == size {
			temp := make([]int, len(transitionList))
			copy(temp, transitionList)
			result = append(result, temp)
			return
		}
		// including the current element
		transitionList = append(transitionList, nums[index])
		recFun(index + 1)
		// not including the current element
		transitionList = transitionList[:len(transitionList)-1]
		recFun(index + 1)
	}
	recFun(0)
	return result
}
