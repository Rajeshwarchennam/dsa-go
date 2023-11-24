package TopViewOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type valueLevelInfo struct {
	val int
	level int
}

func TopViewRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// map with horizontalDistance (-1,0,1,2,3) as keys and value as (node value and level of node)
	hdMap := make(map[int]*valueLevelInfo)
	// min holds the minHd, max holds maxHd
	hdMax, hdMin := 0, 0
	var recFunc func(*TreeNode, int, int) 
	recFunc = func(node *TreeNode, level, hd int) {
		if node == nil {
			return
		}
		if hd > hdMax {
			hdMax = hd
		}
		if hd < hdMin {
			hdMin = hd
		}
		// if for a given hd, a element is already in map, we would replace that with our current node 
		// info if level of current node is less or equals (as it's top view)
		if info, ok := hdMap[hd]; ok {
			if level <= info.level {
				hdMap[hd] = &valueLevelInfo{
					val: node.Val,
					level: level,
				}
			}
		} else {
			hdMap[hd] = &valueLevelInfo{
				val: node.Val,
				level: level,
			}
		}
		recFunc(node.Left, level+1, hd-1)
		recFunc(node.Right, level+1, hd+1)
	}
	recFunc(root, 0, 0)
	res := make([]int, 0, hdMax-hdMin+1)
	// within hdMin to hdMax, we would put values in res array
	for hd := hdMin; hd<=hdMax; hd++ {
		res = append(res, hdMap[hd].val)
	}
	return res
}

func TopViewIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// hdMap carries horzDist as keys and value of nodes as values
	// as we are level order traversing, the very first node we got for a horzDist,
	// will be our top view. We can ignore the upcoming elements which belong to same 
	// horzDist. 
	hdMap, hdMin, hdMax := make(map[int]int), 0, 0
	q := make(queue, 0)
	q.push(&nodeWithHd{
		node: root,
		hd: 0,
	})
	for len(q) != 0 {
		popped := q.pop()
		if _, ok := hdMap[popped.hd]; !ok {
			hdMap[popped.hd] = popped.node.Val
		}
		if popped.hd < hdMin {
			hdMin = popped.hd
		}
		if popped.hd > hdMax {
			hdMax = popped.hd
		}
		if popped.node.Left != nil {
			q.push(&nodeWithHd{
				node: popped.node.Left,
				hd: popped.hd-1,
			})
		}
		if popped.node.Right != nil {
			q.push(&nodeWithHd{
				node: popped.node.Right,
				hd: popped.hd+1,
			})
		}
	}
	res := make([]int, 0, hdMax-hdMin+1)
	for hd := hdMin; hd <= hdMax; hd++ {
		res = append(res, hdMap[hd])
	}
	return res
}

type queue []*nodeWithHd

type nodeWithHd struct {
	node *TreeNode
	hd int
}

func (q *queue) push(x *nodeWithHd) {*q = append(*q, x)}
func (q *queue) pop() *nodeWithHd {
	popped := (*q)[0]
	*q = (*q)[1:len(*q)]
	return popped
}