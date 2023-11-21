package Recursion

func GetPermutationsOfDistinctElements(nums []int) [][]int {
	res := make([][]int, 0, 10)
	size := len(nums)
	transitionList := make([]int, 0, size)
	var recFun func(int, []int)
	recFun = func(indexToFill int, numbersToChooseFrom []int) {
		if indexToFill == size {
			temp := make([]int, size)
			copy(temp, transitionList)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(numbersToChooseFrom); i++ {
			transitionList = append(transitionList, numbersToChooseFrom[i])
			recFun(indexToFill+1, removeithElementFromList(numbersToChooseFrom, i))
			transitionList = transitionList[:len(transitionList)-1]
		}
	}
	recFun(0, nums)
	return res
}
