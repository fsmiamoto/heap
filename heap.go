package heap

import (
	"errors"
)

// Heap is a representation of a binary heap data structure
type Heap struct {
	size     int
	compare  CompareFunc
	elements []interface{}
}

// CompareFunc is a function signature used for comparisons between
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

// New creates a heap using the elements of the slice, with the provided capacity, and
// using the CompareFunc for any comparison between elements
// Since the heap uses a slice underneath, you can specify a initial capacity for it.
// The time complexity of building the heap is O(n), n = len(elements)
func New(elements []interface{}, initialCapacity int, cf CompareFunc) *Heap {
	// Make a copy of the original elements for no bad surprises
	elems := make([]interface{}, len(elements), initialCapacity)
	copy(elems, elements)

	h := &Heap{
		size:     len(elems),
		compare:  cf,
		elements: elems,
	}
	h.heapify()

	return h
}

// Insert adds a new element to the heap
// The time complexity is  O(log(n)), n = # of elements in the heap
func (h *Heap) Insert(x interface{}) {
	h.elements = append(h.elements, x)
	h.size++

	// Fix the heap
	i := h.size - 1
	for i >= 0 && h.compare(h.elements[parent(i)], h.elements[i]) {
		h.elements[parent(i)], h.elements[i] = h.elements[i], h.elements[parent(i)]
		i = parent(i)
	}
}

// Extract returns the element at the root of the heap
// The time complexity is  O(log(n)), n = # of elements in the heap
func (h *Heap) Extract() (interface{}, error) {
	if h.size == 0 {
		return nil, errors.New("heap: empty, no element to extract")
	}

	h.elements[h.size-1], h.elements[0] = h.elements[0], h.elements[h.size-1]
	removedElem := h.elements[h.size-1]

	h.size--

	// Only one node left, no need to fix the heap
	if h.size == 1 {
		return removedElem, nil
	}

	// Fix the heap
	i := 0
	for i < h.size-1 {
		child := h.largerChild(i)

		if h.compare(h.elements[i], h.elements[child]) {
			h.elements[i], h.elements[child] = h.elements[child], h.elements[i]
			if child < h.size/2-1 {
				i = child
			}
		} else {
			break
		}
	}

	return removedElem, nil
}

// IsEmpty indicates if the heap has no elements left
func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

func (h *Heap) largerChild(i int) int {
	left := 2*i + 1
	right := left + 1

	if right > h.size-1 {
		return left
	}

	if h.compare(h.elements[left], h.elements[right]) {
		return right
	}
	return left
}

// heapify makes a heap of the slice in-place
func (h *Heap) heapify() {
	i := h.size/2 - 1
	for i >= 0 {
		left := 2*i + 1
		right := left + 1

		if right > h.size-1 {
			// Look at only the left child
			shouldSwap := h.compare(h.elements[i], h.elements[left])
			if shouldSwap {
				h.elements[i], h.elements[left] = h.elements[left], h.elements[i]
				if left < h.size/2 {
					i = left + 1
				}
			}
		} else {
			// Look at both the left and right child
			rightIsLarger := h.compare(h.elements[left], h.elements[right])
			var compareIndex int

			if rightIsLarger {
				compareIndex = right
			} else {
				compareIndex = left
			}

			shouldSwap := h.compare(h.elements[i], h.elements[compareIndex])
			if shouldSwap {
				h.elements[i], h.elements[compareIndex] = h.elements[compareIndex], h.elements[i]
				if compareIndex < h.size/2 {
					i = compareIndex + 1
				}
			}
		}
		i--
	}
}

func parent(i int) int {
	return (i - 1) / 2
}
