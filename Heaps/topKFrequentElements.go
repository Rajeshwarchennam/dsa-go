package Heaps

import "container/heap"

// link - https://leetcode.com/problems/top-k-frequent-elements/description/

func TopKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, v := range nums {
		countMap[v] += 1
	}
	// creating a minHeap that holds top k frequent nums
	minHeap := make(freqPriorityQueue, 0, k)
	for value, count := range countMap {
		// we perform heapify once we reach size k.
		if minHeap.Len() != k {
			minHeap = append(minHeap, freqElement{
				value: value,
				count: count,
			})
			// if we reach k size, we are initiating the heap
			if minHeap.Len() == k {
				heap.Init(&minHeap)
			}
		} else {
			// if the new element has more freq than head of our minHeap, 
			// we need to remove our head of minHeap and push this element 
			// as our aim is to get top k values. 
			if minHeap[0].count < count {
				minHeap[0] = freqElement{
					value: value,
					count: count,
				}
				// downHeap at index 0
				heap.Fix(&minHeap, 0)
			}
		}
	}
	res := make([]int, 0, k)
	// we are adding the elements in any order as we just needs the top-k
	for _, elem := range minHeap {
		res = append(res, elem.value)
	}
	return res
}

type freqPriorityQueue []freqElement

type freqElement struct {
	value int
	count int
}

func (h freqPriorityQueue) Len() int           { return len(h) }
func (h freqPriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h freqPriorityQueue) Less(i, j int) bool { return h[i].count < h[j].count }

func (h *freqPriorityQueue) Push(x any) { *h = append(*h, x.(freqElement)) }
func (h *freqPriorityQueue) Pop() any {
	old, n := *h, len(*h)
	popped := old[n-1]
	*h = old[:n-1]
	return popped
}