package VerticalOrderTraversal
// link -https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/description/

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type levelMapInfo struct {
	levelMap map[int][]int
	maxLevel int
	minLevel int
}

func VerticalOrderTraversal(root *TreeNode) [][]int {
	// map of horzdistance as key and LevelMaps at that horzDistance as value
	hdMap := make(map[int]*levelMapInfo)
	// to get min and max of horzDist, so we can iterate in that order in hdMap
	hdMax, hdMin := 0, 0
	var recFunc func(*TreeNode, int, int)
	recFunc = func(node *TreeNode, level int, hd int) {
		if node == nil {
			return
		}
		if _, ok := hdMap[hd]; !ok {
			// if that hd is not present, we are creating new value
			hdMap[hd] = &levelMapInfo{
				levelMap: make(map[int][]int),
			}
		}
		lmapInfo := hdMap[hd]
		if _, ok := lmapInfo.levelMap[level]; !ok {
			// if nothing is found in that level yet, we create a new list for that level
			lmapInfo.levelMap[level] = make([]int, 0, 1)
		}
		lmapInfo.levelMap[level] = append(lmapInfo.levelMap[level], node.Val)
		if lmapInfo.maxLevel < level {
			lmapInfo.maxLevel = level
		}
		if lmapInfo.minLevel > level {
			lmapInfo.minLevel = level
		}
		if hd > hdMax {
			hdMax = hd
		}
		if hd < hdMin {
			hdMin = hd
		}
		recFunc(node.Left, level+1, hd-1)
		recFunc(node.Right, level+1, hd+1)
	}
	recFunc(root, 0, 0)
	res := make([][]int, 0, hdMax-hdMin+1)
	for hd := hdMin; hd <= hdMax; hd++ {
		lmapInfo := hdMap[hd]
		// at the current horzDistance, we are adding levelwise
		currentVerticalVals := make([]int, 0, 4)
		for level := lmapInfo.minLevel; level <= lmapInfo.maxLevel; level++ {
			if _, ok := lmapInfo.levelMap[level]; ok {
				sort.Ints(lmapInfo.levelMap[level])
				currentVerticalVals = append(currentVerticalVals, lmapInfo.levelMap[level]...)
			}
		}
		res = append(res, currentVerticalVals)
	}
	return res
}
