package SymmetricTree

// link - https://leetcode.com/problems/symmetric-tree/description/

// TC - O(N)
// SC - O(N)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSymmetric(root *TreeNode) bool {
	// if root is nil, return true
	if root == nil {
		return true
	}
	// this function lets us know, if given trees are mirror images
	var areTreesMirrorImages func(a, b *TreeNode) bool
	areTreesMirrorImages = func(a, b *TreeNode) bool {
		// if either of them is nil, we check if both are equal
		if a == nil || b == nil {
			return a == b
		}
		
		if a.Val != b.Val {
			return false
		}
		// we check if left subtree of a is mirror image of right subtree of b and vice versa
		// if both cases satisfy then, we can say both a and b are mirror images
		return areTreesMirrorImages(a.Left, b.Right) && areTreesMirrorImages(a.Right, b.Left)
	}
	// given tree is symmetric is left and right subtrees are mirror images
	return areTreesMirrorImages(root.Left, root.Right)
}
