package ZigZagTraversal

// link - https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ZigZagTraversal(root *TreeNode) [][]int {
	result := make([][]int, 0, 5)
	if root != nil {
		q := make(queue, 0)
		q.push(root)
		// we use nil as a identifier to represent the end of currentLevel

		q.push(nil)
		// flag to check if we need to add currentLevel values in reverse or normal order
		reverse := false
		// to store currentLevel values
		currentLevel := make([]int, 0)
		for len(q) != 0 {
			popped := q.pop()
			if popped != nil {
				// if popped is a non nil node 
				// we append the values to currentLevel
				// we check left and right and add them to queue.
				// they form next level
				if reverse {
					currentLevel = append([]int{popped.Val}, currentLevel...)
				} else {
					currentLevel = append(currentLevel, popped.Val)	
				}
				if popped.Left != nil {
					q.push(popped.Left)
				}
				if popped.Right != nil {
					q.push(popped.Right)
				}
			} else {
				// if popped is nil, our placeholder to represent the end of current level 
				// we can append the accumulated currentLevel Values to Result
				result = append(result, currentLevel)
				// if popped element is nil and there are no nodes in next level, 
				// there is no need of proceeding further
				if len(q) == 0 {
					break
				}
				// push nil to queue which represent the end of next level
				q.push(nil)
				// recreate currentlevel list to store the nextLevel values
				currentLevel = make([]int, 0)
				// change reverse flag for next level
				reverse = !reverse
			}
		}
	}
	return result
}

type queue []*TreeNode

func (q *queue) push(node *TreeNode) {
	*q = append(*q, node)
}

func (q *queue) pop() *TreeNode {
	popped := (*q)[0]
	*q = (*q)[1:]
	return popped
}

