package FlattenBinaryTree

// link - https://leetcode.com/problems/flatten-binary-tree-to-linked-list/description/

// flatten as preorder traversal

// TC - O(N)
// SC - O(N)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Flatten(root *TreeNode)  {
	// this is prevRecursionNode
	// prevNode is pointed to currentNodes right in everyRecursionCall
	// this updates to currentNode in everyRecursionCall
    var prevNode *TreeNode
	prevNode = nil

	// recursionFunc for dfs that goes right, left, node 

	// we need to form the right most part of this linkedList and then come towards head
	// hence we go in right direction first 
	// we have point our current nodes right to previous recursion call's node
	// that is the reason we are maintaing prevNode 


	var recFunc func(node *TreeNode)
	recFunc = func(node *TreeNode) {
		if node == nil {
			return
		}
		recFunc(node.Right)
		recFunc(node.Left)
		node.Right = prevNode
		node.Left = nil
		prevNode = node
	}
	recFunc(root)
}