package LevelOrderTraversal 

// link - https://leetcode.com/problems/binary-tree-level-order-traversal/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0, 5)
	if root != nil {
		// initialize level
		level := []*TreeNode{root}
		// run loop until we have elements in level
		for len(level) != 0 {
			// this will store nodes of next level
			nextLevel := make([]*TreeNode, 0)
			// values of currentLevel
			currLevelVals := make([]int, 0)
			for _, node := range level {
				currLevelVals = append(currLevelVals, node.Val)
				if node.Left != nil {
					nextLevel = append(nextLevel, node.Left)
				}
				if node.Right != nil {
					nextLevel = append(nextLevel, node.Right)
				}
			}
			result = append(result, currLevelVals)
			// assign nextLevel to CurrentLevel
			level = nextLevel
		}
	}
	return result
}

func LevelOrderII(root *TreeNode) [][]int {
	result := make([][]int, 0, 5)
	if root != nil {
		q := make(queue, 0, 5)
		q.push(root)
		// we use nil as a identifier to represent the end of currentLevel
		q.push(nil)
		// to store currentLevel values
		currentLevel := make([]int, 0)
		for len(q) != 0 {
			popped := q.pop()
			if popped != nil {
				// if popped is a non nil node 
				// we append the values to currentLevel
				// we check left and right and add them to queue.
				// they form next level
				currentLevel = append(currentLevel, popped.Val)
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
				// recreate currentlevel list to store the nextLevel values
				currentLevel = make([]int, 0)
				// push nil to queue which represent the end of next level
				q.push(nil)
			}
		}
	}
	return result
}

type queue []*TreeNode

func (q *queue) pop() *TreeNode {
	popped := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return popped
}

func (q *queue) push(node *TreeNode) {
	*q = append(*q, node)
}
