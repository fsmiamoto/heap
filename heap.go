package heap

import "errors"

type Heap struct {
	size     int
	capacity int
	compare  CompareFunc
	elements []interface{}
}

type CompareFunc func(node, child interface{}) bool

func New(elements []interface{}, capacity int, cf CompareFunc) *Heap {
	// Make a copy of the original elements for no bad surprises
	elems := make([]interface{}, capacity)
	copy(elems, elements)
	heapify(elems, cf)
	return &Heap{
		size:     len(elems),
		capacity: capacity,
		compare:  cf,
		elements: elems,
	}
}

func (h *Heap) Insert(x interface{}) error {
	if h.size == h.capacity {
		return errors.New("heap has no capacity")
	}

	h.size++
	h.elements[h.size-1] = x

	// Fix the heap
	i := h.size - 1
	for i >= 0 && h.compare(h.elements[parent(i)], h.elements[i]) {
		h.elements[parent(i)], h.elements[i] = h.elements[i], h.elements[parent(i)]
		i = parent(i)
	}

	return nil
}

func (h *Heap) DeleteRoot() interface{} {
	return nil
}

// heapify makes a heap of the slice in-place
// TODO: review this logic
func heapify(elems []interface{}, compare CompareFunc) {
	i := len(elems)/2 - 1
	for i >= 0 {
		left := 2*i + 1
		right := 2*i + 2

		if right > len(elems)-1 {
			// Look at only the left child
			shouldSwap := compare(elems[i], elems[left])
			if shouldSwap {
				elems[i], elems[left] = elems[left], elems[i]
				if left < len(elems)/2 {
					i = left + 1
				}
			}
			continue
		}

		rightIsLarger := compare(elems[left], elems[right])
		var compareIndex int

		if rightIsLarger {
			compareIndex = right
		} else {
			compareIndex = left
		}

		shouldSwap := compare(elems[i], elems[compareIndex])
		if shouldSwap {
			elems[i], elems[compareIndex] = elems[compareIndex], elems[i]
			if compareIndex < len(elems)/2 {
				i = compareIndex + 1
			}
			continue
		}
		i--
	}
}

func parent(i int) int {
	return (i - 1) / 2
}
