package SearchInBST

// link - https://leetcode.com/problems/search-in-a-binary-search-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	// if we found the val at current node or reached beyond leaf and couldn't find the val, we can return node. 
	if root == nil || root.Val == val {return root}
	if val > root.Val {
		return searchBST(root.Right, val)
	}
	return searchBST(root.Left, val)
} 