package LeftOrRightViewOfBinaryTree

// link - https://practice.geeksforgeeks.org/problems/left-view-of-binary-tree/1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LeftViewOfBinaryTree(root *TreeNode) []int {
	res := make([]int, 0, 5)
	var recFunc func(*TreeNode, int)
	// we would move in node -> left -> right direction
	// we have parameter level which corresponds to level of currentNode
	// whenever we get element for the first time in that level, we add to our result
	recFunc = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level == len(res) {
			res = append(res, node.Val)
		}
		recFunc(node.Left, level+1)
		recFunc(node.Right, level+1)
	}
	recFunc(root, 0)
	return res
}

func RightViewOfBinaryTree(root *TreeNode) []int {
	res := make([]int, 0, 5)
	var recFunc func(*TreeNode, int)
	// we would move in node -> right -> left direction
	// we have parameter level which corresponds to level of currentNode
	// whenever we get element for the first time in that level, we add to our result
	recFunc = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level == len(res) {
			res = append(res, node.Val)
		}
		recFunc(node.Right, level+1)
		recFunc(node.Left, level+1)
	}
	recFunc(root, 0)
	return res
}