package InorderTraversal

// link - https://leetcode.com/problems/binary-tree-inorder-traversal/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversalRecursive(root *TreeNode) []int {
	// res to store inorder traversal
	res := make([]int, 0, 5)
	var recFunc func(*TreeNode)
	recFunc = func (node *TreeNode)  {
		if node == nil {
			return
		}
		// left -> curr -> right
		recFunc(node.Left)
		res = append(res, node.Val)
		recFunc(node.Right)
	}
	// calling the recFunc from root
	recFunc(root)
	return res
}

func InorderTraversalIterative(root *TreeNode) []int {
	// stack to mimic recursive stack calls
	s := make(stack, 0, 5)
	// we start from root
	currNode := root
	res := make([]int, 0, 5)
	// starting from currNode, in the left direction, push everything to stack
	// when all the left elements are over, we can pop from stack (and add that to our result) and pick right element 
	// of popped element and again continue in left direction
	for currNode != nil || len(s) != 0 {
		if currNode != nil {
			s.Push(currNode)
			currNode = currNode.Left
		} else {
			popped := s.Pop()
			res = append(res, popped.Val)
			currNode = popped.Right
		}
	}
	return res
}

type stack []*TreeNode

func (s *stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *stack) Pop() *TreeNode {
	popped :=  (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped
}