package DiameterOfBinaryTree

// link - https://leetcode.com/problems/diameter-of-binary-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DiameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	// this recFunc is to find height of tree with root as 'node'
	var  recFunc func(*TreeNode) int
	recFunc = func(node *TreeNode) int{
		if node == nil {
			return 0
		}
		// we find leftHeight and rightHeight of tree with 'node' as root
		leftHeight := recFunc(node.Left)
		rightHeight := recFunc(node.Right)
		// findig diameter passing through this node as a curve 
		// longest path passing through this node as a curve is leftHeight+rightHeight of tree with 'node' as root
		// we are finding max of all the diameters 
		diameter = max(diameter, leftHeight+rightHeight)
		return 1 + max(leftHeight, rightHeight)
	}
	recFunc(root)
	return diameter
}

func max(a, b int) int {
	if a > b {return a}
	return b
}