package KthSmallestElement

// link - https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// kth smallest element (1-indexed)
func KthSmallestElement(root *TreeNode, k int) int {
	ans := 0

	var recFun func(node *TreeNode)
	// inorderTraversal with small variation
	// we keep on reducting k-value, once we got kth smallest element, we no longer call any recursions and return to parent
	recFun = func(node *TreeNode) {

		if node == nil {
			return
		}

		recFun(node.Left)

		k--
		if k == 0 {
			ans = node.Val
			return
		}
		
		if k < 0 {
			return
		}

		recFun(node.Right)
	}

	recFun(root)
	return ans
}