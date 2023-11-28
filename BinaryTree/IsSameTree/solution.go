package IsSameTree

// link -https://leetcode.com/problems/same-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p, q *TreeNode) bool {
	// if both p and q are nil, we can return true
	if p == nil && q == nil {
		return true
	}
	// if only one of them is nil, then they are not same trees
	if p == nil {
		return false
	}
	if q == nil {
		return false
	}
	// if values of p and q are different, then they aren't same
	if p.Val != q.Val {
		return false
	}
	// if left subtrees are same and right subtrees are same, 
	// then we can state entire p, q trees are same
	if IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right) {
		return true
	}
	return false
}