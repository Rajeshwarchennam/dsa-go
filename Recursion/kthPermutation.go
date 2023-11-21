package Recursion

import "strconv"

func GetKthPermutation(n int, k int) string {
	numbers := make([]int, 0, n)

	fact := 1
	// forming (n-1)! for fixing first spot
	for i := 1; i < n; i++ {
		fact *= i
		numbers = append(numbers, i)
	}
	// adding the nth element to numbers
	numbers = append(numbers, n)
	ans := ""
	// we are reducing k by 1 as we want to start k from 0 instead of 1.
	// for example n=3, numbers = [1,2,3].To get 0th,1st element starts with 1nums[0], 2nd, 3rd elemnt starts with 2nums[1],
	// hence we can get that index by k/fact. if k=3, we can get index by 3/2! = 1. nums[1] = 2 _ _.
	k = k - 1
	for {
		// as told above k/fact index is current slot
		ans = ans + strconv.Itoa(numbers[k/fact])
		// as that element is used, we can remove from numbers.
		numbers = removeithElementFromList(numbers, k/fact)
		// if we used all numbers, we reached our solution
		if len(numbers) == 0 {
			break
		}
		// once a slot is fixed, after that k would be k%fact
		// example if n=3, k= 2(0th index). first slot is k/fact = 2/(3-1)! = 1. 2 _ _. starting with 2, it would be 0th element.
		// so k becomes k%fact = 2%2! = 0
		k = k % fact
		fact = fact / len(numbers)
	}
	return ans
}

func removeithElementFromList(nums []int, index int) []int {
	var res []int
	n := len(nums)
	if n == 1 {
		res = []int{}
	} else if index == 0 {
		res = append([]int{}, nums[index+1:]...)
	} else if index == n-1 {
		res = append([]int{},nums[:index]...)
	} else {
		res = append(append([]int{},nums[:index]...), nums[index+1:]...)
	}
	return res
}
