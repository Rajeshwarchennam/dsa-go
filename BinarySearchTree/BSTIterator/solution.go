package BSTIterator

// link - https://leetcode.com/problems/binary-search-tree-iterator/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
    s *stack
}


func Constructor(root *TreeNode) BSTIterator {
    stack := make(stack, 0, 5)
	pushAllLeft(root, &stack)
	return BSTIterator{
		s: &stack,
	}
}


func (this *BSTIterator) Next() int {
    popped := this.s.pop()
	pushAllLeft(popped.Right, this.s)
	return popped.Val
}


func (this *BSTIterator) HasNext() bool {
    return len(*this.s) != 0
}

func pushAllLeft(node *TreeNode, s *stack) {
	for node != nil {
		s.push(node)
		node = node.Left
	}
}

type stack []*TreeNode

func (s *stack) push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *stack) pop() *TreeNode {
	n := len(*s)
	popped := (*s)[n-1]
	*s = (*s)[:n-1]
	return popped
}