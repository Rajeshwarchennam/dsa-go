package Heaps

import "container/heap"
// link - https://leetcode.com/problems/kth-largest-element-in-an-array/description/
func FindKthLargestElementUsingMaxHeap(nums []int, k int) int {
	n := len(nums)
	// maxHeap whose aim is to hold small n-k+1 values. Therefore, 0th index would be the
	// kth largest element
	maxHeap := make(maxHeapsInt, n-k+1)
	copy(maxHeap, nums)
	// initiate a maxHeap with first n-k+1 values
	heap.Init(&maxHeap)
	for i := n-k+1; i < n; i++ {
		// for rest of elements, we would check if it is smaller than 
		// highest element in maxHeap. We would replace if it is (as we aim for holding 
		// smaller values in our maxHeap)
		if nums[i] < maxHeap[0] {
			maxHeap[0] = nums[i]
			// fix the heap after replacing with new element
			heap.Fix(&maxHeap, 0)
		}
	}
	// 0th index would be the highest of all n-k+1 small values of nums
	// i.e kth largest element
	return maxHeap[0]
}

type maxHeapsInt []int

func (h maxHeapsInt) Len() int           { return len(h) }
func (h maxHeapsInt) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h maxHeapsInt) Less(i, j int) bool { return h[i] > h[j] }

func (h *maxHeapsInt) Push(x any) { *h = append(*h, x.(int)) }
func (h *maxHeapsInt) Pop() any {
	old := *h
	n := len(old)
	popped := old[n-1]
	*h = old[:n-1]
	return popped
}