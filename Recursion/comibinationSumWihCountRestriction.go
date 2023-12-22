package Recursion

import "sort"

// link - https://leetcode.com/problems/combination-sum-ii/


// given list of (non-distinct)nums and a target find all unique combination of nums which sum to target. 
// each num in nums may only be used once. 

func CombinationSumWithCountRestriction(nums []int, target int) [][]int {
	size := len(nums)
	res := make([][]int, 0, size)
	transitionList := make([]int, 0, 5)
	var recFunc func(index, currTarget int)
	// this recFunc helps in filling the current Position in transitionList
	recFunc = func(index, currTarget int) {
		// if target is reached, we can append tl in result and return
		if currTarget == 0 {
			temp := make([]int, len(transitionList))
			copy(temp, transitionList)
			res = append(res,  temp)
			return
		}
		// for filling current position we can select any element from index to size-1 indices.
		// if we created a combination by putting an element at currentPosition, we need to avoid 
		// putting the same duplicate element again in that position and form a combination.
		for i := index; i < size; i++ {
			// avoiding currentIndex, we check for all duplicates and avoid adding them
			// in combination for current position
			if i != index && nums[i] == nums[i-1] {
				continue
			}
			// if nums[i] is beyond our currTarget, we can come out of loop as including and beyond that element,
			// they won't form combination. 
			if nums[i] > currTarget {
				break
			}
			transitionList = append(transitionList, nums[i])
			// recFunc for next position in transitionList.
			// we would start picking elements from i+1 index.
			recFunc(i+1, currTarget-nums[i])
			// we remove the added element in transitionList as we would try combination
			// with other element at the position
			transitionList = transitionList[:len(transitionList)-1]
		}
	}
	// sort nums for applying above recFunc
	sort.Ints(nums)
	recFunc(0, target)
	return res
}