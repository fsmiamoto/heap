package heap_test

import (
	"fmt"
	"log"

	"github.com/fsmiamoto/heap"
)

// This example shows a use in the Heapsort algorithm
func Example_heapSort() {
	values := []interface{}{40, 30, 50, 100, 15}
	h := heap.New(values, len(values), heap.MinInt)

	var sorted []int
	for !h.IsEmpty() {
		v, err := h.Extract()
		if err != nil {
			log.Fatal(err)
		}
		sorted = append(sorted, v.(int))
	}
	fmt.Println(sorted)
	// Output:
	// [15 30 40 50 100]
}
