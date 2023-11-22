package PreOrderTraversal
// link - https://leetcode.com/problems/binary-tree-preorder-traversal/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversalRecursive(root *TreeNode) []int {
    res := make([]int, 0, 5)
	var recFunc func(*TreeNode)
	recFunc = func (node *TreeNode)  {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		recFunc(node.Left)
		recFunc(node.Right)
	}
	recFunc(root)
	return res
}



func PreOrderTraversalIterative(root *TreeNode) []int {
	res := make([]int, 0, 5)
	s := make(stack, 0, 10)
	currNode := root 
	// node -> left -> right
	// first add currNode val to res. Then move in the left direction and keep on putting in stack
	// (add val to res before putting in stack)
	// when we hit nil in that direction, we pop from stack and pick right of popped element and 
	// again move in left direction
	for currNode != nil || len(s) != 0 {
		if currNode != nil {
			res = append(res, currNode.Val)
			s.push(currNode)
			currNode = currNode.Left
		} else {
			currNode = s.pop().Right
		}
	}
	return res
}

type stack []*TreeNode

func (s *stack) push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *stack) pop() *TreeNode{
	n := len(*s)
	popped := (*s)[n-1]
	*s = (*s)[:n-1]
	return popped
}