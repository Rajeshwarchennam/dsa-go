package BalancedBinaryTree

// link - https://leetcode.com/problems/balanced-binary-tree/description/
import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	// this recFunc would return height of tree if it is balanced or -1 if it is not a balanced tree
	var recFunc func(*TreeNode) int
	recFunc = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		leftHeight := recFunc(tn.Left)
		// if we get -1, we understand it's not a balanced tree
		// hence don't find height of other subtrees
		// just return -1 (which states entire tree is not balanced)
		if leftHeight == -1 {
			return -1
		}
		rightHeight := recFunc(tn.Right)
		if rightHeight == -1 {
			return -1
		}
		// if diffrence b/w leftHeight and rightHeight is greater than 1, we can say it as not balanced.
		// hence need not compute for other subtrees
		if math.Abs(float64(leftHeight-rightHeight)) > 1 {
			return -1
		}
		return int(math.Max(float64(leftHeight), float64(rightHeight)))+1
	}
	return recFunc(root) != -1 
}