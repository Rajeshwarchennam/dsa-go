package Heaps

import "container/heap"

// link - https://leetcode.com/problems/find-median-from-data-stream/description/

type MinOrMaxHeap struct {
	Values   []int
	LessFunc func(int, int) bool
}

func (h MinOrMaxHeap) Less(i, j int) bool { return h.LessFunc(h.Values[i], h.Values[j]) }
func (h MinOrMaxHeap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h MinOrMaxHeap) Len() int           { return len(h.Values) }
func (h *MinOrMaxHeap) Push(x any)        { h.Values = append(h.Values, x.(int)) }
func (h *MinOrMaxHeap) Pop() (popped any) {
	h.Values, popped = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return popped
}

func NewMinMaxHeap(lessFunc func(int, int) bool) *MinOrMaxHeap {
	return &MinOrMaxHeap{
		LessFunc: lessFunc,
	}
}

type MedianFinder struct {
	minHeap *MinOrMaxHeap // to hold 2nd half of stream (including the median if stream is of odd count)
	maxHeap *MinOrMaxHeap // to hold 1st half of stream
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: NewMinMaxHeap(func(a, b int) bool { return a < b }),
		maxHeap: NewMinMaxHeap(func(a, b int) bool { return a > b }),
	}
}

func (mf *MedianFinder) AddNum(num int) {
	// if stream count is even, after adding the new element,
	// the extra element should flow to minHeap (2nd half).
	// hence we put new num in maxHeap(1st half) and remove the max among that half
	// and put it in 2nd half i.e minHeap
	if (mf.minHeap.Len()+mf.maxHeap.Len())%2 == 0 {
		heap.Push(mf.maxHeap, num)
		heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))
	} else {
		// if stream count is odd, after adding the new element, the extra element should
		// flow to first half (maxHeap). Hence we put new num in minHeap(2nd half) and pop 
		// the min of Second Half(minHeap) and put it in firstHalf (maxHeap)
		heap.Push(mf.minHeap, num)
		heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if (mf.maxHeap.Len()+mf.minHeap.Len())%2 == 0 {
		return float64(mf.minHeap.Values[0]+mf.maxHeap.Values[0])/2.0
	}
	return float64(mf.minHeap.Values[0])
}