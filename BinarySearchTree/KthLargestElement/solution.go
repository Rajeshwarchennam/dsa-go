package KthLargestElement

// link - https://practice.geeksforgeeks.org/problems/kth-largest-element-in-bst/1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func KthLargestElement(root *TreeNode, k int) int {
	ans := 0

	var recFunc func(node *TreeNode)
	// reverse inorderTraversal with small variation
	// we keep on reducting k-value, once we got kth largest element, we no longer call any recursions and return to parent
	recFunc = func (node *TreeNode)  {

		if node == nil {
			return
		}

		recFunc(node.Right)

		k--
		if k == 0 {
			ans = node.Val
			return
		}
		
		if k < 0 {
			return
		}

		recFunc(node.Left)	
	}

	return ans
}