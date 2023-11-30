package MaximumPathSum

// link - https://leetcode.com/problems/binary-tree-maximum-path-sum/description/

// TC - O(N)
// SC - O(N)
import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaximumPathSum(root *TreeNode) int {
	// this stores the maximumPathSum at any node 
	// with that node as curved point of path
	maximumPathSum := math.MinInt64
	// this recFunc would return maximum sum starting from that node 
	// either in left direction of node or right (if any of them or non negative)
	var recFunc func(node *TreeNode) int 
	recFunc = func(node *TreeNode) int {
		// if given node is nil, starting with that node sum would be zero
		if node == nil {
			return 0 
		}
		// find leftPathSum and rightPathSum
		leftPathMaxSum := recFunc(node.Left)
		rightPathMaxSum := recFunc(node.Right)
		// we find maximumPathSum as an intermediate step in finding maximumSum starting with given node.
		// for our node to be used as curved point of path, it needs to be added to leftPathSum and rightPathSum(if they are non-negative)
		// We are adding node.val though it is negative, as we want atleast one element in the path
		maximumPathSum = max(maximumPathSum, node.Val+max(leftPathMaxSum, 0)+max(rightPathMaxSum, 0))
		// we return the maximumSum starting with the given node, 
		// either by taking left or rightPath which is max (if any of them is non-negative)
		return node.Val + max(max(leftPathMaxSum, rightPathMaxSum), 0)
	}
	recFunc(root)
	return maximumPathSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}