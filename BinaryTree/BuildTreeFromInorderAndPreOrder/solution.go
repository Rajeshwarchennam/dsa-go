package BuildTreeFromInorderAndPreOrder

// link - https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

// TC - O(N)
// SC - O(N)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(preorder, inorder []int) *TreeNode {
	n := len(preorder)
	// if there are no elements in both orders, we can return nil
	if n == 0 {
		return nil
	}
	// start value of preorder traversal is root
	rootVal := preorder[0]
	// finding the index of root in inorder traversal
	rootIndexInInorder := 0
	for i := 0; i<n; i++ {
		if inorder[i] == rootVal {
			rootIndexInInorder = i
			break
		}
	}
	// elements to the left of root in inorder forms left Subtree and right forms right subtree.
	// In preorder, the count of elements after first element(root) that forms left subtree preorder is 
	// count of elements that form the Inorder of leftsubtree. 
	// similarly with right subtree
	root := &TreeNode{
		Val: rootVal,
		Left: BuildTree(preorder[1:1+rootIndexInInorder], inorder[:rootIndexInInorder]),
		Right: BuildTree(preorder[1+rootIndexInInorder:], inorder[rootIndexInInorder+1:]),
	}
	return root
}