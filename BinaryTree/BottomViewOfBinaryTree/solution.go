package BottomViewOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type valueLevelInfo struct {
	val int
	level int
}

// for a given horizontal distance, our aim to choose element with highest level
func BottomViewRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	// map with horizontalDistance (-1,0,1,2,3) as keys and value as (node value and level of node)
	horzDistanceMap := make(map[int]*valueLevelInfo)
	// min holds the minHd, max holds maxHd
	horzDistanceMin, horzDistanceMax := 0, 0

	var recFunc func(*TreeNode, int, int)
	recFunc = func(node *TreeNode, level, hd int) {
		if node == nil {
			return
		}
		if hd > horzDistanceMax {
			horzDistanceMax = hd
		} 
		if hd < horzDistanceMin {
			horzDistanceMin = hd
		}
		// if for a given hd, a element is already in map, we would replace that with our current node 
		// info if level of current node is greater or equals
		if info, ok := horzDistanceMap[hd]; ok {
			if level >= info.level {
				horzDistanceMap[hd] = &valueLevelInfo{
					val: node.Val,
					level: level,
				}
			}
		} else {
			// if given hd is not in map yet, we would add our node info against that hd in map
			horzDistanceMap[hd] = &valueLevelInfo{
				val: node.Val,
				level: level,
			}
		}
		recFunc(node.Left, level+1, hd-1)
		recFunc(node.Right, level+1, hd+1)
	}

	recFunc(root, 0, 0)
	res := make([]int, 0, horzDistanceMax-horzDistanceMin+1)
	// within hdMin to hdMax, we would put values in res array
	for hd := horzDistanceMin; hd <= horzDistanceMax; hd++ {
		res = append(res, horzDistanceMap[hd].val)
	}
	return res
}

func BottomViewIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// horizontalDistanceMap would carry horzDist as keys and value of nodes as values
	// as we are following level order traversal, we need not keep track of levels.
	// whenever we are getting new value for a hd, we would only pick that value as 
	// that belongs >= previous node level which present in that hd.
	horzDistanceMap, horzDistanceMin, horzDistanceMax := make(map[int]int), 0, 0
	q := make(queue, 0)
	// start with root and hd as 0
	q.Push(&nodeWithHd{
		node: root,
		hd:0,
	})
	for len(q) != 0 {
		popped := q.Pop()
		// as told we can keep adding latest elements without checking level 
		// as this is level order traversal
		horzDistanceMap[popped.hd] = popped.node.Val
		if popped.hd > horzDistanceMax {horzDistanceMax = popped.hd}
		if popped.hd < horzDistanceMin {horzDistanceMin = popped.hd}
		if popped.node.Left != nil {
			q.Push(&nodeWithHd{
				node: popped.node.Left,
				hd: popped.hd-1,
			})
		}
		if popped.node.Right != nil {
			q.Push(&nodeWithHd{
				node: popped.node.Right,
				hd: popped.hd+1,
			})
		}
	}
	res := make([]int, 0, horzDistanceMax-horzDistanceMin+1)
	// within hdMin to hdMax, we would put values in res array
	for hd := horzDistanceMin; hd <= horzDistanceMax; hd++ {
		res = append(res, horzDistanceMap[hd])
	}
	return res	
}

type nodeWithHd struct {
	node *TreeNode
	hd int
}

type queue []*nodeWithHd
func (q *queue) Push(x *nodeWithHd) {*q = append(*q, x)}
func (q *queue) Pop() *nodeWithHd {
	popped := (*q)[0]
	*q = (*q)[1:len(*q)]
	return popped
}