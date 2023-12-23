package ConstructBstFromPreorder

import (
	"math"
)

// link - https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	size := len(preorder)
	// this pointer to index int helps us to traverse entire preorder traversal through all recursion calls
	var index *int = new(int)
	*index = 0
	// this function helps in building bst
	// when value at current index if less than or equals to upperbound, we can create a node out of that value 
	// and return that node to parent. Increment index value stating we have built tree till that index before.
	// if current index node is greater than upperBound, we can return nil to parent call stating we need to put that 
	// index value else where. 
	// if we have traversed all the values in preorder, we can return nil. 
	var build func( upperBound int) *TreeNode
	build = func(upperBound int) *TreeNode{
		if *index == size || preorder[*index] > upperBound {return nil}
		root := &TreeNode{Val : preorder[*index]}
		// increment index stating we have seen the currentIndex 
		*index ++
		// for left of the node, node.Val be the upperBound
		root.Left = build(root.Val)
		// for right of the node, upperBound of node itself is the upperBound. 
		// We need not care about lowerBound because all those elements will any way move to left subtree as we are 
		// calling it first. 
		root.Right = build(upperBound)
		return root
	}
	return build(math.MaxInt64)
}