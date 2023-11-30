package BuildTreeFromInorderAndPostorder

// link - https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/

// TC - O(N)
// SC - O(N)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(inorder, postorder []int) *TreeNode {
	n := len(inorder)
	// if there are no elements in both orders, we can return nil
	if n == 0 {
		return nil
	}
	//end value of postorder is root
	rootVal := postorder[n-1]
	// finding the index of root in inorder traversal
	rootIndexInInorder := 0 
	for i := 0; i < n; i++ {
		if inorder[i] == rootVal {
			rootIndexInInorder = i
			break
		}
	}
	// elements to the left of root in inorder forms left Subtree and right forms right subtree.
	// In postorder, the count of first elements that forms left subtree is 
	// count of elements that form the Inorder of leftsubtree. 
	// similarly with right subtree (next n elements that forms right subtree is similar to that count in inorder)
	root := &TreeNode{
		Val: rootVal,
		Left: BuildTree(inorder[:rootIndexInInorder], postorder[:rootIndexInInorder]),
		Right: BuildTree(inorder[rootIndexInInorder+1:], postorder[rootIndexInInorder:n]),
	}
	return root
}