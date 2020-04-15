package heap

type Heap struct {
	size     int
	capacity int
	elements []interface{}
}

type CompareFunc func(node, child interface{}) bool

func New(elements []interface{}, capacity int, c CompareFunc) *Heap {
	h := &Heap{
		size:     len(elements),
		capacity: capacity,
	}

	elems := make([]interface{}, capacity)
	copy(elems, elements)

	i := len(elems)/2 - 1
	for i >= 0 {
		left := 2*i + 1
		right := 2*i + 2

		switch {
		case h.validIndex(left) && !h.validIndex(right):
			shouldSwap := c(elems[i], elems[left])
			if shouldSwap {
				elems[i], elems[left] = elems[left], elems[i]
				if left < len(elems)/2 {
					i = left + 1
				}
			}
		case h.validIndex(left) && h.validIndex(right):
			rightIsLarger := c(elems[left], elems[right])
			var compareIndex int

			if rightIsLarger {
				compareIndex = right
			} else {
				compareIndex = left
			}

			shouldSwap := c(elems[i], elems[compareIndex])
			if shouldSwap {
				elems[i], elems[compareIndex] = elems[compareIndex], elems[i]
				if compareIndex < len(elems)/2 {
					i = compareIndex + 1
				}
			}
		}
		i--
	}

	h.elements = elems
	return h
}

func (h *Heap) Insert(x interface{}) {
	return
}

func (h *Heap) DeleteRoot() interface{} {
	return nil
}

func (h *Heap) validIndex(i int) bool {
	return i < h.size
}

func parent(i int) int {
	return (i - 1) / 2
}
