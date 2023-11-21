package Recursion

import "sort"

func CombinationSumWithCountRestriction(nums []int, target int) [][]int {
	size := len(nums)
	res := make([][]int, 0, size)
	transitionList := make([]int, 0, 5)
	var recFunc func(index, currTarget int)
	recFunc = func(index, currTarget int) {
		if currTarget == 0 {
			temp := make([]int, len(transitionList))
			copy(temp, transitionList)
			res = append(res,  temp)
			return
		}
		for i := index; i < size; i++ {
			if i != index && nums[i] == nums[i-1] {
				continue
			}
			if nums[i] > currTarget {
				break
			}
			transitionList = append(transitionList, nums[i])
			recFunc(i+1, currTarget-nums[i])
			transitionList = transitionList[:len(transitionList)-1]
		}
	}
	sort.Ints(nums)
	recFunc(0, target)
	return res
}