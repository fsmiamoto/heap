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
	h.buildHeap()

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
	h.heapify(0)

	return removedElem, nil
}

// IsEmpty indicates if the heap has no elements left
func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

// buildHeap makes a heap of the slice in-place
func (h *Heap) buildHeap() {
	for i := h.size/2 - 1; i >= 0; i-- {
		child := h.largerChild(i)

		shouldSwap := h.compare(h.elements[i], h.elements[child])
		if !shouldSwap {
			continue
		}

		h.elements[i], h.elements[child] = h.elements[child], h.elements[i]
		if child < h.size/2 {
			i = child + 1
		}
	}
}

// heapify fixes the heap property at root index i.
// This function assumes the subtrees are already heapified
func (h *Heap) heapify(i int) {
	child := h.largerChild(i)

	if child > h.size-1 {
		return
	}

	shouldSwap := h.compare(h.elements[i], h.elements[child])
	if !shouldSwap {
		return
	}

	h.elements[i], h.elements[child] = h.elements[child], h.elements[i]
	h.heapify(child)
}

// largerChild returns the index of the larger child, as defined with the
// provided CompareFunc
func (h *Heap) largerChild(i int) int {
	left := 2*i + 1
	right := left + 1

	if right > h.size-1 {
		// There's not a right child
		return left
	}

	if h.compare(h.elements[left], h.elements[right]) {
		return right
	} else {
		return left
	}
}

//parent retuns the index of the parent of the given child index
func parent(child int) int {
	return (child - 1) / 2
}
