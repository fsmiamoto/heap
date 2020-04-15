package heap

import (
	"reflect"
	"testing"
)

// CompareFunc's
func max(node, child interface{}) bool {
	return child.(int) > node.(int)
}

func min(node, child interface{}) bool {
	return child.(int) < node.(int)
}

func TestHeap(t *testing.T) {
	tests := []struct {
		name   string
		cf     CompareFunc
		input  []interface{}
		expect []interface{}
	}{
		{"MaxHeap", max, []interface{}{10, 20, 15, 12, 40, 25, 18}, []interface{}{40, 20, 25, 12, 10, 15, 18}},
		{"MinHeap", min, []interface{}{8, 12, 9, 7, 22, 3, 26, 14, 11, 15, 22}, []interface{}{3, 7, 8, 11, 15, 9, 26, 14, 12, 22, 22}},
	}

	for _, tt := range tests {
		heap := New(tt.input, len(tt.input), tt.cf)

		if len(heap.elements) != len(tt.expect) {
			t.Errorf("Heap has the wrong number of elements")
		}

		equal := reflect.DeepEqual(heap.elements, tt.expect)
		if !equal {
			t.Errorf("Expected %v but got %v", tt.expect, heap.elements)
		}
	}
}
