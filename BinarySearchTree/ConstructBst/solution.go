package ConstructBst

// link - https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/

// given sorted array, convert to bst (height balanced)


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	// if we reach beyond leaf, we can return nil
	if len(nums) == 0 {
		return nil
	}
	// to achieve height balanced bst, we need to put mid element as root at any instance
	mid := len(nums)/2
	return &TreeNode{
		Val: nums[mid],
		Left: sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}