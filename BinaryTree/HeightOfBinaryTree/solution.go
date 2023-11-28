package HeightOfBinaryTree

// link - https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HeightOfBinaryTreeRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(HeightOfBinaryTreeRecursion(root.Left), HeightOfBinaryTreeRecursion(root.Right)) + 1
}

func HeightOfBinaryTreeIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := list.New()
	q.PushBack(&qElement{root, 1})
	height := 0
	for q.Len() != 0 {
		popped := q.Remove(q.Front()).(*qElement)
		height = max(height, popped.height)
		if popped.node.Left != nil {
			q.PushBack(&qElement{popped.node.Left, popped.height+1})
		}
		if popped.node.Right != nil {
			q.PushBack(&qElement{popped.node.Right, popped.height+1})
		}
	}
	return height
}

type qElement struct {
	node   *TreeNode
	height int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}