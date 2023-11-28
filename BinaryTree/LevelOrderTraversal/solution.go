package LevelOrderTraversal 

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