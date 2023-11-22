package Heaps

import "container/heap"
// link - https://www.geeksforgeeks.org/merge-k-sorted-arrays-set-2-different-sized-arrays/

func MergeKSortedArraysOfDifferentSize(arrays [][]int, k int) []int {
	// use minHeap to store k values
	minHeap := make(intMinHeap, 0, k)
	// add the min of all the arrays
	for i := 0; i < k; i++ {
		minHeap = append(minHeap, valueWithIndex{
			value: arrays[i][0],
			index: []int{i, 0},
		})
	}
	heap.Init(&minHeap)
	res := make([]int, 0, 10)
	// while the minHeap is not empty
	for minHeap.Len() != 0 {
		// we would pop the minimumelement in heap and add it to result
		// and we will check which array is giving this min element and we 
		// will add the next element from same array if available (as possibly that can be next possible 
		// element)
		minElement := heap.Pop(&minHeap).(valueWithIndex)
		res = append(res, minElement.value)
		// checking if next element is present in the array from which the current 
		// minelement has come. 
		if len(arrays[minElement.index[0]]) > minElement.index[1]+1 {
			heap.Push(&minHeap, valueWithIndex{
				value: arrays[minElement.index[0]][minElement.index[1]+1],
				index: []int{minElement.index[0], minElement.index[1]+1},
			})
		}
	}
	return res
}

type valueWithIndex struct {
	value int
	// this holds index info. if nth array and ith index in that array means [n,i]
	index []int
}

type intMinHeap []valueWithIndex

func (h intMinHeap) Len() int           { return len(h) }
func (h intMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h intMinHeap) Less(i, j int) bool { return h[i].value < h[j].value }

func (h *intMinHeap) Push(x any) { *h = append(*h, x.(valueWithIndex)) }
func (h *intMinHeap) Pop() any {
	old := *h
	n := len(old)
	popped := old[n-1]
	*h = old[:n-1]
	return popped
}