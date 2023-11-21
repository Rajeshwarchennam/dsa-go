package Heaps

import (
	"container/heap"
	"sort"
)

// link - https://www.interviewbit.com/problems/maximum-sum-combinations/

func MaxCombinationSums(a, b []int, k int) []int {
	n := len(a)
	// sort a, b in asc order.
	sort.Ints(a)
	sort.Ints(b)
	// this map helps to check if some index pair is already added to heap
	// for example [i-1, j] results in adding [i-1,j-1] and [i-2, j]
	// and [i, j-1] results in adding [i-1, j-1] and [i, j-2]
	// if we dont track the visited map, [i-1, j-1] pair sum will be added twice
	visitedMap := make(map[indexPair]bool)	

	maxHeap := make(maxSumHeap, 0)
	// appending the highest sum element
	maxHeap = append(maxHeap, sumIndexesObj{
		indexes: indexPair{
			i: n-1,
			j: n-1,
		},
		sum: a[n-1]+b[n-1],
	})
	// marking (n-1,n-1) as visited
	visitedMap[indexPair{n-1, n-1}] = true
	heap.Init(&maxHeap)

	res := make([]int, 0, k)

	// iterating k times to pop k elements
	for i := 0; i < k; i++ {
		// popping out the maxElement
		maxElem := heap.Pop(&maxHeap).(sumIndexesObj)
		// adding to result
		res = append(res, maxElem.sum)

		// finding the nextPossible pair that needs to added if not visited
		nextIndexPair := indexPair{maxElem.indexes.i, maxElem.indexes.j-1}
		if !visitedMap[nextIndexPair] {
			heap.Push(&maxHeap, sumIndexesObj{
				indexes: nextIndexPair, 
				sum: a[nextIndexPair.i]+b[nextIndexPair.j],
			})
			visitedMap[nextIndexPair] = true
		}
		// second possible pair
		nextIndexPair = indexPair{maxElem.indexes.i-1, maxElem.indexes.j}
		if !visitedMap[nextIndexPair] {
			heap.Push(&maxHeap, sumIndexesObj{
				indexes: nextIndexPair, 
				sum: a[nextIndexPair.i]+b[nextIndexPair.j],
			})
			visitedMap[nextIndexPair] = true
		}
	}
	return res
}

type sumIndexesObj struct {
	indexes indexPair
	sum int
}

type indexPair struct {
	i int
	j int
}

type maxSumHeap []sumIndexesObj

func (h maxSumHeap) Less(i, j int) bool  {return h[i].sum > h[j].sum}
func (h maxSumHeap) Len() int {return len(h)}
func (h maxSumHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func (h *maxSumHeap) Push(x any) {*h = append(*h, x.(sumIndexesObj))}
func (h *maxSumHeap) Pop() any {
	old := *h
	n := len(old)
	popped := old[n-1]
	*h = old[:n-1]
	return popped
}