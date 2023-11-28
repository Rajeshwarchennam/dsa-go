package LowestCommonAncestor

// link - https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// if root is nil, we return nil
	// if we find currentNode as p or q, we return it
	// we won't check in it's subtree. if we find node as p, by returning p 
	// our solution works if q is present in node's sub tree or other part of the tree (as we know p and q exists in tree)
	if root == nil || root == p || root == q {
		return root
	}
	// find the leftLca and right lca
	leftLca := LowestCommonAncestor(root.Left, p, q)

	rightLca := LowestCommonAncestor(root.Right, p, q)

	// if we got one from left subtree and other from right subtree, 
	// we can state current node is lca
	if leftLca != nil && rightLca != nil {
		return root
	}
	// if leftLca is not nil and rightLca is nil; 
	// then we can say leftLca is Lca as we only propagate lca values to parent. 
	if leftLca != nil {
		return leftLca
	}
	return rightLca
}