package RunningFromBothEnds

// link-  https://leetcode.com/problems/next-permutation/description/

// For example, for arr = [1,2,3], the following are all the permutations of arr: [1,2,3], [1,3,2], [2, 1, 3], [2, 3, 1], [3,1,2], [3,2,1].

func nextPermutation(nums []int)  {
    n  :=  len(nums)

	// this is used to find the index after which the nums sequence drops
	// example 1354, we are finding index of 3. Coz after 3, 54 is non increasing sequence
	// technically we are finding the element for which post the element the sequence has maxed out
	
	indexBeforeDescent := -1 

	for i := n-2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			indexBeforeDescent = i
			break
		}
	}

	// if we found the indexBeforeDescent, we are finding the value just greater than the value at indexBeforeDescent
	// to form the next sequence
	if indexBeforeDescent != -1 {
		for i := n-1; i > indexBeforeDescent; i-- {
			if nums[i] > nums[indexBeforeDescent] {
				// once we find the value just greater than value at indexBeforeDescent we swap values
				nums[i], nums[indexBeforeDescent] = nums[indexBeforeDescent], nums[i]
				break
 			}
		}
	}

	// we reverse the values of non increasing sequence now to non decreasing sequence (as that produces the first permutation 
	// by placing value that is just greater than the value which was at indexBeforeDescent) 
	start, end := indexBeforeDescent+1, n-1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start += 1
		end -= 1
	}
}

func reverseNums(nums []int, start, end int) {
    for i,j := start, end; i < j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
} 