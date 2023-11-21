package Heaps

// implementing heap from scratch
type Item interface {
	Less (than Item) bool 
}

type Heap struct{
	items []Item
}

func NewHeap(items []Item) *Heap {
	h := &Heap{items: items}
	heapify(h)
	return h
}

func heapify(h *Heap) {
	for i := len(h.items)/2-1; i >= 0; i-- {
		downHeap(h, i)
	}
}

func downHeap(h *Heap, i int) {
	for {
		leftChild, rightChild := 2*i+1, 2*i+2
		// nextChild to which if at all this downHeap needs to be implemented.
		var nextChild int 
		if leftChild < len(h.items) && rightChild < len(h.items) {
			if h.items[leftChild].Less(h.items[rightChild]) {
				nextChild = leftChild
			} else {
				nextChild = rightChild
			}
		} else if leftChild < len(h.items) {
			nextChild = leftChild
		} else {
			// no children means, we are already satisfying heap rule
			break
		}
		if h.items[nextChild].Less(h.items[i]) {
			h.items[nextChild], h.items[i] = h.items[i], h.items[nextChild]
			i = nextChild
		} else {
			// if parent is lesser than smaller of the two childs, heap property is satisfied
			break
		}
	}
}

func (h *Heap) Push(item Item) {
	h.items = append(h.items, item)
	upHeap(h, len(h.items)-1)
}

func upHeap(h *Heap, i int) {
	for {
		parent := (i-1)/2
		if parent < 0 || h.items[parent].Less(h.items[i]) {
			// if heap rule already satisfies, then break
			break 
		}
		h.items[i], h.items[parent] = h.items[parent], h.items[i]
		i = parent
	}
}

func (h *Heap) Pop() Item {
	if len(h.items) == 0 {
		return nil
	}
	popped := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	downHeap(h, 0)
	return popped
}


// If you want to implement above heap for int items, Implement Item interface like below

//type intItem int

// func (i intItem) Less(than Heaps.Item) bool {
// 	if t, ok := than.(intItem); ok {
// 		return i < t
// 	}
// 	return false
// }

// items := []Heaps.Item{intItem(1), intItem(10), intItem(-1), intItem(-100), intItem(-11)}
// h :=  Heaps.NewHeap(items)
// h.Push(intItem(11))