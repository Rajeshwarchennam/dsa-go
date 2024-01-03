package MaxSumBSTInBT

import "math"

// link - https://leetcode.com/problems/maximum-sum-bst-in-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxSumBST(root *TreeNode) int {

	result := 0
	var recFunc func(node *TreeNode) sumMinMax
	
	// post order traversal, where we find the sum, min, max of bsts starting with currentNode
	// The underlying tree is bst, if currentNode's val is greater than max of left subtree
	// and less than min of right subtree

	recFunc = func(node *TreeNode) sumMinMax {
		// if we reached nil nodes, we need to return sumMinMax such that,
		// if this nil appears towards left, the max should always be less than leaf node val
		// if this nil appears towards right, the min should always be greater than leaf node val
		// sum will anyway be 0
		if node == nil {
			return sumMinMax{
				sum: 0,
				min: math.MaxInt64,
				max: math.MinInt64,
			}
		}
		leftSumMinMax := recFunc(node.Left)
		rightSumMinMax := recFunc(node.Right)
		// if currentNode value lies between left subtree max value and right subtree min value
		// we can state from node, it is a bst.
		if leftSumMinMax.max < node.Val && node.Val < rightSumMinMax.min {
			sum := node.Val + leftSumMinMax.sum + rightSumMinMax.sum
			result = maxInt(result, sum)
			return sumMinMax{
				sum: sum,
				min: minInt(node.Val, leftSumMinMax.min),
				max: maxInt(node.Val, rightSumMinMax.max),
			}
		}
		// if it is not a bst, 
		// we need to make sure the parent of this node should always fail to be a bst
		// hence it's min is set to most min int (this fails the bst logic when current node appears to be right subtree of parent node)
		// it's max is set to mst max int (this fails the bst logic when current ndodes appears to the left of parent node)
		return sumMinMax{
			sum: 0,
			min: math.MinInt64,
			max: math.MaxInt64,
		}
	}

	recFunc(root)
	return result
}


// this holds sum, min, max values of the bst starting from that node
type sumMinMax struct {
	sum int
	min int
	max int
}

func maxInt(a,b int) int {
	if a > b {return a}
	return b
}

func minInt(a, b int) int {
	if a < b {return a}
	return b
}