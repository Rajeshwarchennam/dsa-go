package ValidateBST

import "math"

// link - https://leetcode.com/problems/validate-binary-search-tree/description/


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}


// isValid function checks if the current node value lies between min and max bounds.
// if a node is nil, we can say it be valid subtree/tree of bst
// we also check if left and right subtrees are valid bsts by sending their min and max bounds
func isValid(node *TreeNode, min, max int) bool {
	if node == nil { return true}
	if (node.Val >= max || node.Val <= min) {return false}
	return isValid(node.Left, min, node.Val) && isValid(node.Right, node.Val, max)
}