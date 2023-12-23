package PopulateNextRightPointer

// link - https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

// Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root != nil {
		q := make(queue, 0, 5)
		// previous node in the same level
		// when at node in level order traversal, we will use this prevNode's Next to point to current node

		var prevNode *Node
		prevNode = nil

		q.Push(root)
		// this nil will be used as placeholder to state the end of current level 
		q.Push(nil)

		for len(q) != 0 {
			popped := q.Pop()
			if popped != nil {
				if prevNode != nil {
					prevNode.Next = popped
				}
				if popped.Left != nil {
					q.Push(popped.Left)
				}
				if popped.Right != nil {
					q.Push(popped.Right)
				}
				prevNode = popped
			} else {
				// if no elements in next level, we can break
				if len(q) == 0 {
					break
				}
				// reInitialize prevNode
				prevNode = nil
				// add placeHolder to state next level end
				q.Push(nil)
			}
		}
	}
	return root
}

type queue []*Node

func (q *queue) Push(node *Node) { *q = append(*q, node) }
func (q *queue) Pop() *Node {
	popped := (*q)[0]
	*q = (*q)[1:]
	return popped
}
