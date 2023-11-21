package Heaps

import "container/heap"
// link - https://leetcode.com/problems/kth-largest-element-in-an-array/description/
func FindKthLargestElementUsingMinHeap(nums []int, k int) int {
	// we are forming a minHeap with size k. 
	// Aim is to hold k largest elements in this heap. As it's minHeap, Oth index is kth Largest element. 
	minHeap := make(minHeapInts, k)
	copy(minHeap, nums)
	// initiating heap with first k elements
	heap.Init(&minHeap)
	for i := k; i<len(nums); i++ {
		// we would check if the new element needs to be put in our heap. 
		// our aim is to keep hold of only k large elements
		// hence we would remove smallest element in heap, when we found a large value
		if nums[i] > minHeap[0] {
			minHeap[0] = nums[i]
			// this would fix heap at 0th index
			heap.Fix(&minHeap, 0)
		}
	}
	// returning the smallest element in heap (which holds k large elements), i.e Kth largest element.
	return minHeap[0]
}

type minHeapInts []int

func (h minHeapInts) Len() int           { return len(h) }
func (h minHeapInts) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h minHeapInts) Less(i, j int) bool { return h[i] < h[j] }

func (h *minHeapInts) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeapInts) Pop() any {
	old := *h
	n := len(old)
	popped := old[n-1]
	*h = old[:n-1]
	return popped
}