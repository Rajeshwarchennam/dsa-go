package TwoSumProblem 

// link - https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


// SC - O(H)
// TC - O(N)

func TwoSumProblem(root *TreeNode, target int) bool {

	if root == nil {return false}
	
	leftPointer, rightPointer := NewBstNextIteraor(root), NewBstPrevIteraor(root)

	// similar to two sum problem in sorted arrays. We move left and right pointers 
	// based on their sum and comparing with target
	l, r := leftPointer.next(), rightPointer.prev()


	// considering all vals in tree are unique. 
	for l < r {
		if l + r == target {return true}
		if l + r < target {
			l = leftPointer.next()
		} else {
			r = rightPointer.prev()
		}
	}
	return false

}


// In this iterator stack, we put all the elements towards left.
// when we pop an element by calling next, we see if that node has any right.
// if yes, we put all the left elements of that right node in stack. 

type BSTNextIterator struct {
	s *stack
}

func (i *BSTNextIterator) hasNext() bool {
	return len(*i.s) != 0
}

func (i *BSTNextIterator) next() int {
	popped := i.s.pop()
	if popped.Right != nil {
		pushAllLeft(popped.Right, i.s)
	}
	return popped.Val
}

type BSTPrevIterator struct {
	s *stack
}

func (i *BSTPrevIterator) hasPrev() bool {
	return len(*i.s) != 0
}

func (i *BSTPrevIterator) prev() int {
	popped := i.s.pop()
	if popped.Left != nil {
		pushAllRight(popped.Left, i.s)
	}
	return popped.Val
}

func NewBstNextIteraor(root *TreeNode) *BSTNextIterator {
	stack := make(stack, 0, 5)
	pushAllLeft(root, &stack)
	return &BSTNextIterator{
		s: &stack,
	}
}

func NewBstPrevIteraor(root *TreeNode) *BSTPrevIterator {
	stack := make(stack, 0, 5)
	pushAllRight(root, &stack)
	return &BSTPrevIterator{
		s: &stack,
	}
}




func pushAllLeft(node *TreeNode, s *stack) {
	for node != nil {
		s.push(node)
		node = node.Left
	}
}

func pushAllRight(node *TreeNode, s *stack) {
	for node != nil {
		s.push(node)
		node = node.Right
	}
}

type stack []*TreeNode

func (s *stack) push(tn *TreeNode) {
	*s = append(*s, tn)
}

func (s *stack) pop() *TreeNode {
	n := len(*s)
	popped := (*s)[n-1]
	*s = (*s)[:n-1]
	return popped
}