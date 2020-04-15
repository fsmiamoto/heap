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

	t.Run("Test Insert", func(t *testing.T) {
		tests := []struct {
			name       string
			element    interface{}
			cf         CompareFunc
			shouldFail bool
			heap       *Heap
			expect     *Heap
		}{
			{"Insert on MaxHeap", 60, max, false,
				&Heap{
					size:     7,
					capacity: 9,
					compare:  max,
					elements: []interface{}{50, 30, 20, 15, 10, 8, 16, 0, 0},
				},
				&Heap{
					size:     8,
					capacity: 9,
					compare:  max,
					elements: []interface{}{60, 50, 20, 30, 10, 8, 16, 15, 0},
				},
			},
			{"Heap already full", 60, max, true,
				&Heap{
					size:     9,
					capacity: 9,
					compare:  max,
					elements: []interface{}{50, 30, 20, 15, 10, 8, 16, 9, 8},
				},
				&Heap{
					size:     9,
					capacity: 9,
					compare:  max,
					elements: []interface{}{50, 30, 20, 15, 10, 8, 16, 9, 8},
				},
			},
		}

		for _, tt := range tests {
			err := tt.heap.Insert(tt.element)
			assertError(t, tt.shouldFail, err)
			assertEqualHeap(t, tt.heap, tt.expect)

		}
	})
}

func assertError(t *testing.T, shouldFail bool, err error) {
	t.Helper()
	if shouldFail && err == nil {
		t.Errorf("a error was expected but got none")
	}

	if !shouldFail && err != nil {
		t.Errorf("no error was expected but got %v", err)
	}
}

func assertEqualHeap(t *testing.T, got, want *Heap) {
	t.Helper()
	if got.size != want.size {
		t.Errorf("heaps have different sizes, got %d want %d", got.size, want.size)
	}
	if got.capacity != want.capacity {
		t.Errorf("heaps have different capacities, got %d want %d", got.capacity, want.capacity)
	}
	if !reflect.DeepEqual(got.elements, want.elements) {
		t.Errorf("heaps have different elements, got %v want %v", got.elements, want.elements)
	}
}
