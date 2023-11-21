package Heaps
// implementing the Heap interfaces. Lib support
type IntMinHeap []int

func (h IntMinHeap) Len() int {
	return len(h)
}

func (h IntMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	// need to remove last element. Not the first as we did in scratch implementation. This would be taken care by lib
	popped := old[n-1]
	*h = old[:n-1]
	return popped	
}

// initiate items like below to use heap on our items. 
//items := &Heaps.IntMinHeap{23,1,2,1,2,-10,-9}
//heap.Init(items)