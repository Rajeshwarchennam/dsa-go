package ChildrenSumProperty


// link - https://takeuforward.org/data-structure/check-for-children-sum-property-in-a-binary-tree/

// TC - O(N)
// SC - O(N)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ChangeTree(root *TreeNode) {
	if root == nil {
		return
	}
	childSum := 0 
	if root.Left != nil {
		childSum += root.Left.Val
	}
	if root.Right != nil  {
		childSum += root.Right.Val
	}


	if childSum < root.Val {
	
		// if childSUm is less than parent Val, make children value equal to parent value
		// we are increasing the children value here so we make sure we never fall short
		// hence while coming back up we can just add children value and replace parent
		if root.Left != nil {
			root.Left.Val = root.Val
		}
		if root.Right != nil {
			root.Right.Val = root.Val
		}
	}

	// call changeTree for both left and right trees
	ChangeTree(root.Left)
	ChangeTree(root.Right)

	// now we can sum up children values and replace the parent value 
	// if both the children are not nil
	childSum = 0 
	if root.Left != nil {
		childSum += root.Left.Val
	}
	if root.Right != nil {
		childSum += root.Right.Val
	}
	if root.Left != nil || root.Right != nil {
		root.Val = childSum
	}
}