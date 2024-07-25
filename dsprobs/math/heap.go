package math

import "fmt"

type Heap struct {
	items []int
}

func (h *Heap) AddToHeap(item int) {
	h.items = append(h.items, item)
	h.heapifyUp(len(h.items) - 1)
	fmt.Println(h.items)
}

func (h *Heap) heapifyUp(index int) {
	for index > 0 {
		if h.items[index] > h.items[h.parent(index)] {
			h.swap(index, h.parent(index))
			index = h.parent(index)
		} else {
			break
		}
	}
}

func (h *Heap) RemoveElementFromHeap() int {
	elementExtracted := h.items[0]
	lastAvailableIndex := len(h.items) - 1
	h.items[0] = h.items[lastAvailableIndex]
	h.items = h.items[:lastAvailableIndex]
	h.heapifyDown(0)
	return elementExtracted
}

func (h *Heap) heapifyDown(index int) {
	l, r := h.left(index), h.right(index)
	lastIndex := len(h.items) - 1
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if l < r {
			childToCompare = r
		} else {
			childToCompare = l
		}
		if h.items[index] < h.items[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
		} else {
			break
		}
	}
}

func (h *Heap) left(index int) int {
	return 2*index + 1
}

func (h *Heap) right(index int) int {
	return 2*index + 2
}

func (h *Heap) parent(index int) int {
	return (index - 1) / 2
}

func (h *Heap) swap(index1, index2 int) {
	h.items[index1], h.items[index2] = h.items[index2], h.items[index1]
}
