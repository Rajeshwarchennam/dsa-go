package RunningFromBothEnds

// link - https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/ 

// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

func twoSum(numbers []int, target int) []int {
    i, j := 0, len(numbers)-1
    for i < j {
        temp := numbers[i] + numbers[j]
        if temp > target {
            j -= 1
        } else if temp < target {
            i += 1
        } else {
            return []int{i+1,j+1}
        }
    }
    return []int{0,0}
}
