package heap

import (
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	tests := []struct {
		name   string
		cf     CompareFunc
		input  []interface{}
		expect []interface{}
	}{
		{"MaxHeap #1", MaxInt,
			[]interface{}{10, 20, 15, 12, 40, 25, 18},
			[]interface{}{40, 20, 25, 12, 10, 15, 18}},
		{"MaxHeap #2", MaxInt,
			[]interface{}{100, 88, 24, 43, 35, 7, 97, 39, 46, 38, 15, 53, 65, 93, 87, 84, 59, 52, 24, 53},
			[]interface{}{100, 88, 97, 84, 53, 65, 93, 59, 52, 38, 15, 53, 7, 24, 87, 39, 43, 46, 24, 35}},
		{"MaxHeap #3", MaxInt,
			[]interface{}{10, 20},
			[]interface{}{20, 10}},
		{"MaxHeap #4", MaxInt,
			[]interface{}{5},
			[]interface{}{5}},
		{"MaxHeap #5", MaxInt,
			[]interface{}{},
			[]interface{}{}},
		{"MinHeap #1", MinInt,
			[]interface{}{8, 12, 9, 7, 22, 3, 26, 14, 11, 15, 22},
			[]interface{}{3, 7, 8, 11, 15, 9, 26, 14, 12, 22, 22}},
		{"MinHeap #2", MinInt,
			[]interface{}{40, 30, 50, 100, 15},
			[]interface{}{15, 30, 50, 100, 40}},
		{"MinHeap #3", MinInt,
			[]interface{}{15, 7, 96, 38, 54, 52, 69, 25, 28, 46, 18, 39, 6, 7, 29, 88, 40, 37, 67, 22},
			[]interface{}{6, 7, 7, 25, 18, 39, 15, 38, 28, 22, 54, 96, 52, 69, 29, 88, 40, 37, 67, 46}},
		{"MinHeap #3", MinInt,
			[]interface{}{10, 20},
			[]interface{}{10, 20}},
		{"MinHeap #4", MinInt,
			[]interface{}{5},
			[]interface{}{5}},
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
			{"Insert on MaxHeap", 60, MaxInt, false,
				&Heap{
					size:     7,
					capacity: 9,
					compare:  MaxInt,
					elements: []interface{}{50, 30, 20, 15, 10, 8, 16, 0, 0},
				},
				&Heap{
					size:     8,
					capacity: 9,
					compare:  MaxInt,
					elements: []interface{}{60, 50, 20, 30, 10, 8, 16, 15, 0},
				},
			},
			{"Heap already full", 60, MaxInt, true,
				&Heap{
					size:     9,
					capacity: 9,
					compare:  MaxInt,
					elements: []interface{}{50, 30, 20, 15, 10, 8, 16, 9, 8},
				},
				&Heap{
					size:     9,
					capacity: 9,
					compare:  MaxInt,
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

	t.Run("Test Extract", func(t *testing.T) {
		tests := []struct {
			name             string
			cf               CompareFunc
			shouldFail       bool
			extractedElement interface{}
			heap             *Heap
			expect           *Heap
		}{
			{"Extract on MaxHeap", MaxInt, false, 50,
				&Heap{
					size:     7,
					capacity: 9,
					compare:  MaxInt,
					elements: []interface{}{50, 30, 20, 15, 10, 8, 16, 0, 0},
				},
				&Heap{
					size:     6,
					capacity: 9,
					compare:  MaxInt,
					elements: []interface{}{30, 16, 20, 15, 10, 8, 50, 0, 0},
				},
			},
			{"Extract on small MaxHeap", MaxInt, false, 50,
				&Heap{
					size:     2,
					capacity: 2,
					compare:  MaxInt,
					elements: []interface{}{50, 30},
				},
				&Heap{
					size:     1,
					capacity: 2,
					compare:  MaxInt,
					elements: []interface{}{30, 50},
				},
			},
			{"Extract on MinHeap", MinInt, false, 3,
				&Heap{
					size:     11,
					capacity: 11,
					compare:  MinInt,
					elements: []interface{}{3, 7, 8, 11, 15, 9, 26, 14, 12, 22, 22},
				},
				&Heap{
					size:     10,
					capacity: 11,
					compare:  MinInt,
					elements: []interface{}{7, 11, 8, 12, 15, 9, 26, 14, 22, 22, 3},
				},
			},
			{"Empty heap", MinInt, true, nil,
				&Heap{
					size:     0,
					capacity: 3,
					compare:  MinInt,
					elements: []interface{}{0, 0, 0},
				},
				&Heap{
					size:     0,
					capacity: 3,
					compare:  MinInt,
					elements: []interface{}{0, 0, 0},
				},
			},
		}

		for _, tt := range tests {
			element, err := tt.heap.Extract()
			assertEqual(t, element, tt.extractedElement)
			assertError(t, tt.shouldFail, err)
			assertEqualHeap(t, tt.heap, tt.expect)

		}
	})
}

func assertEqual(t *testing.T, got, expect interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Element is not the expected one, got %v but expected %v", got, expect)
	}
}

func assertError(t *testing.T, shouldFail bool, err error) {
	t.Helper()
	if shouldFail && err == nil {
		t.Errorf("A error was expected but got none")
	}

	if !shouldFail && err != nil {
		t.Errorf("No error was expected but got %v", err)
	}
}

func assertEqualHeap(t *testing.T, got, want *Heap) {
	t.Helper()
	if got.size != want.size {
		t.Errorf("Heaps have different sizes, got %d want %d", got.size, want.size)
	}
	if got.capacity != want.capacity {
		t.Errorf("Heaps have different capacities, got %d want %d", got.capacity, want.capacity)
	}
	if !reflect.DeepEqual(got.elements, want.elements) {
		t.Errorf("Heaps have different elements, got %v want %v", got.elements, want.elements)
	}
}
