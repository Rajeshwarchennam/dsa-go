package Recursion

import "strings"

func SolveNKings(n int) [][]string {
	res := make([][]string, 0, 5)

	transitionList := make([]int, 0, n)
	var recFun func(int)
	recFun = func(currRow int) {
		// if we reached n, we can append our answer
		if currRow == n {
			temp := make([]string, 0, len(transitionList))
			// for each val in transitionList we are making a string like "..K." (val=2, n=4)
			for _, val := range transitionList {
				currRowString := strings.Repeat(".", val) + "K" + strings.Repeat(".", n-val-1)
				temp = append(temp, currRowString)
			}
			res = append(res, temp)
			return
		}
		for i := 0; i<n; i++ {
			// checking for previous row element and avoiding in contact to that
			if len(transitionList) != 0 && (i == transitionList[len(transitionList)-1] || i == transitionList[len(transitionList)-1]+1 || i == transitionList[len(transitionList)-1]-1) {
				continue
			}
			// considering all the contactless possiblities
			transitionList = append(transitionList, i)
			recFun(currRow+1)
			transitionList = transitionList[:len(transitionList)-1]
		}
	}
	recFun(0)

	return res
}