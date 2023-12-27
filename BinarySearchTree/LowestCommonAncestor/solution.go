package LowestCommonAncestor

// link - https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	currentNode := root
	
	// find the min and max values of p.Val and q.Val. These help in 
	// iterating through tree
	minVal, maxVal := min(p.Val, q.Val), max(p.Val, q.Val)
	for {
		// if currenode value lies between minVal and maxVal, 
		// it acts as pivot between these nodes (considering a node to be descendent to itself)
		if currentNode.Val >= minVal && currentNode.Val <= maxVal {
			// only break method for the loop, as it is guaranteed that p and q exist 
			return currentNode
		} else if currentNode.Val < minVal {
			currentNode = currentNode.Right
		} else {
			currentNode = currentNode.Left
		}
	}
}

func min(a, b int) int {
	if a < b {return a}
	return b
}

func max(a, b int) int {
	if a > b {return a}
	return b
}