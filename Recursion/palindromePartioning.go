package Recursion

// link - https://leetcode.com/problems/palindrome-partitioning/description/

// given a string s, return all possible palindrome partitionings

func PalindromePartioning(s string) [][]string {
	size := len(s)
	res := make([][]string, 0, 10)
	transitionList := make([]string, 0, 5)
	var recFunc func(index int)
	recFunc = func(index int) {
		// this means we have completed the last partioning including string[size-1] element also
		// hence add that list of palindrome substrings to res
		if index == size {
			temp := make([]string, len(transitionList))
			copy(temp, transitionList)
			res = append(res, temp)
			return
		}
		// this for loop means starting from index, we have make a partion after i index such that s[index:i+1] is a palindrome
		for i := index; i < size; i++ {
			if !isPalindrome(s, index, i) {
				continue
			}
			// if we got s[index:i+1] as palindrome, we can add that to our transitionList and check for another partion from i+1 element. Last partion
			// was made right after i index element
			transitionList = append(transitionList, s[index:i+1])
			recFunc(i + 1)
			transitionList = transitionList[:len(transitionList)-1]
		}
	}
	recFunc(0)
	return res
}

func isPalindrome(s string, start, end int) bool {
	for start <= end {
		if s[start] != s[end] {
			return false
		}
		start += 1
		end -= 1
	}
	return true
}