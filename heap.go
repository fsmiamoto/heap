package heap

import "errors"

// Heap is a representation of a binary heap data structure
type Heap struct {
	size     int
	capacity int
	compare  CompareFunc
	elements []interface{}
}

// CompareFunc is a function signature used for comparions between
// a node and it's children, returning true if the two should be swapped
type CompareFunc func(node, child interface{}) bool

// MaxInt is a CompareFunc for a MaxHeap of integers
func MaxInt(node, child interface{}) bool {
	return child.(int) > node.(int)
}

// MinInt is a CompareFunc for a MinHeap of integers
func MinInt(node, child interface{}) bool {
	return child.(int) < node.(int)
}

// New creates a heap using the elements of the slice with the provided capacity and
// using the CompareFunc for any comparison, therefore you can a
// MaxHeap or a MinHeap just by changing the function
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

// Insert adds a new element to the heap
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

// Extract returns the element at the root of the heap
func (h *Heap) Extract() interface{} {
	panic("not implemented")
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
