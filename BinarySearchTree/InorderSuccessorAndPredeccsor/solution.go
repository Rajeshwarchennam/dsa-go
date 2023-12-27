package InorderSuccessorAndPredeccsor

// link - https://practice.geeksforgeeks.org/problems/predecessor-and-successor/1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


// Return Successor node Of targetNode whose value is just greater than targetNode.Val
func InorderSuccessor(root, targetNode *TreeNode) *TreeNode {
	var successorNode *TreeNode = nil 

	currentNode := root 

	// until we reach a leaf 
	for currentNode != nil {
		// if targetVal is greater than/ equals current Node value, 
		// we iterate in right direction to find successor. 
		if targetNode.Val >= currentNode.Val {
			currentNode = currentNode.Right
		} else {
			// if targetVal is less than currentNode val, we can state this can be a possible successor 
			// we are sure this is lesser than our previous found successor as by the above iterations, 
			// we are getting closer to the targetVal
			successorNode = currentNode
			currentNode = currentNode.Left
		}
	}

	return successorNode
}

func InorderPredeccsor(root, targetNode *TreeNode) *TreeNode {
	var predeccsorNode *TreeNode = nil 

	currentNode := root

	for currentNode != nil {
		if targetNode.Val <= currentNode.Val {
			currentNode = currentNode.Left
		} else {
			predeccsorNode = currentNode
			currentNode = currentNode.Right
		}
	}
	return predeccsorNode
}