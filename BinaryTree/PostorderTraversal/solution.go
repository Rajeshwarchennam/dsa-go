package PostorderTraversal 

// link - https://leetcode.com/problems/binary-tree-postorder-traversal/description/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// left -> right -> root
func PostorderTraversalRecursive(root *TreeNode) []int {
	res := make([]int, 0, 5)
	var recFunc func(*TreeNode)
	recFunc = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		recFunc(tn.Left)
		recFunc(tn.Right)
		res = append(res, tn.Val)
	}
	recFunc(root)
	return res
}


//               1 
//          2          3
//      4      5    6     7

// post order traversal of above is 4, 5, 2, 6, 7, 3, 1
// reverse of that is 1, 3, 7, 6, 2, 5, 4 which looks like preOrderTraversal but root->right->left
// hence we can implement similar to preorder traversal but towards right direction
// and reversing the solution 

func PostOrderTraversalIterative(root *TreeNode) []int {
	res := make([]int, 0, 5)
	s := make(stack, 0, 5)
	currentNode := root
	// following node -> right -> left
	for currentNode != nil || len(s) != 0 {
		if currentNode != nil {
			res = append(res, currentNode.Val)
			s.Push(currentNode)
			currentNode = currentNode.Right
		} else {
			currentNode = s.Pop().Left
		}
	}
	reverseList(res)
	return res
}

func reverseList(list []int) {
	i, j := 0, len(list)-1
	for i <= j {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
}

type stack []*TreeNode

func (s *stack) Push(node *TreeNode) {
	*s = append(*s, node)
}
func (s *stack) Pop() *TreeNode {
	n := len(*s)
	popped := (*s)[n-1]
	*s = (*s)[:n-1]
	return popped
}