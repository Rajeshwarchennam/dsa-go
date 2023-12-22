package Recursion


// link - https://leetcode.com/problems/combination-sum/description/
// Time complexity - O(2^target*k) where k is average length of each sumSet


// Problem - given distint nums and target, return a list of all unique combinations of nums where sum is target
// we can choose same num any number of times

func CombinationSum(nums []int, target int) [][]int {
	size := len(nums)
	res := make([][]int, 0, size)
	transitionList := make([]int, 0, 5)
	var recFun func(index int, currTarget int)
	recFun = func (index int, currTarget int)  {
		// reached end of nums
		if index == size {
			// if only currTarget reduced to 0 then add our solution to res
			if currTarget == 0 {
				temp := make([]int, len(transitionList))
				copy(temp, transitionList)
				res = append(res, temp)
			}
			return
		}
		// avoid adding element if it is greater than currTarget
		if nums[index] <= currTarget {
			transitionList = append(transitionList, nums[index])
			recFun(index, currTarget-nums[index])
			// need to removet the element after above recursion as parent function doesn't include the added element
			transitionList = transitionList[:len(transitionList)-1]
		}
		// non pick current element case
		recFun(index+1, currTarget)
	}
	recFun(0, target)
	return res
}