package RunningFromBothEnds

// link - https://leetcode.com/problems/trapping-rain-water/description/ 

// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

func trap(height []int) int {
    max_left, max_right, left, right, result := 0, 0, 0, len(height)-1, 0
    for left <= right {
        // aim of left and right pointers is to move towards max in heights
        
        // as we know we are moving left pointer towards right as right pointer value is larger.  
        // we can state that all the lefts viewed till now (including max_left) is 
        // smaller that height[right]. Therefore we can have safewall all the time towards right
        // we can just check with left wall that is max_left.
        if height[left] <= height[right] {
            max_left = max(max_left, height[left])
            result += max_left - height[left]
            left ++
        } else {
            max_right = max(max_right, height[right])
            result += max_right - height[right]
            right --
        }
    }
    return result
}

func max(a, b int) int {
    if a > b {return a}
    return b
}

func min(a,b int) int {
    if a < b {return a}
    return b
}
