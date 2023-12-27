package FloorInBST

// link - https://www.codingninjas.com/studio/problems/floor-from-bst_920457

// given a value X, find the greatest value in BST which is smaller than or equals to  X. X is not smaller than smallest node of BST.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FloorInBST(root *TreeNode, x int) int {

	floor := 0

	currentNode := root

	for currentNode != nil {
		if currentNode.Val == x {
			floor = x
			break
		} else if currentNode.Val > x {

			currentNode = currentNode.Left
		} else {
			floor = currentNode.Val
			currentNode = currentNode.Right
		}
	}
	return floor
}
