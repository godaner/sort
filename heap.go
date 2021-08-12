package sort

type Heap struct {
	len int
}

func (h *Heap) Sort(src Sortable) (dst Sortable) {
	h.len = src.Len()
	src = h.buildHeap(src)
	for h.len > 0 {
		src.Swap(0, h.len-1)
		h.len--
		src = h.heapify(src, 0)
	}
	return src
}

func (h *Heap) buildHeap(src Sortable) (dst Sortable) {
	for i := src.Len() / 2; i >= 0; i-- {
		src = h.heapify(src, i)
	}
	return src
}

// heapify
//  从上至下递归调整堆大小
func (h *Heap) heapify(src Sortable, i int) (dst Sortable) {
	maxI, left, right := i, i*2+1, i*2+2
	if left < h.len && src.Less(maxI, left) {
		maxI = left
	}
	if right < h.len && src.Less(maxI, right) {
		maxI = right
	}
	if maxI != i {
		src.Swap(maxI, i)
		src = h.heapify(src, maxI)
	}
	return src
}
