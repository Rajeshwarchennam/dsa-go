package PathFromRootToNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func PathFromRootToNode(root *TreeNode, nodeVal int) []int {
	res := make([]int, 0, 5)
	var recFunc func(*TreeNode) bool 
	recFunc = func(tn *TreeNode) bool {
		// if we reach nil, we can say we didn't find our node in that path
		// so we return false
		if tn == nil {
			return false
		}
		// we will put our current node val in path and proceed to check if can find path by including this
		res = append(res, tn.Val)
		// if we reached req node we can return true
		if tn.Val == nodeVal {
			return true
		}
		// if either of left or right paths got the req value, we can return true
		if recFunc(tn.Left) || recFunc(tn.Right) {
			return true
		}
		// if none of left and right paths get us to req value, we can remove curr element from our path 
		// and return false stating to parent call that this child node doesn't route to required node
		res = res[:len(res)-1]
		return false
	}
	recFunc(root)
	return res
}