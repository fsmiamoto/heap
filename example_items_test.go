package heap_test

import (
	"fmt"
	"log"

	"github.com/fsmiamoto/heap"
)

// Item is an example defined type
type Item struct {
	key int
}

// MaxItem is an example CompareFunc that builds a MaxHeap using the key property
func MaxItem(node, child interface{}) bool {
	return child.(Item).key > node.(Item).key
}

// This example shows how to implement your own CompareFunc and the use on
// a defined type
func Example_items() {
	values := []interface{}{
		Item{key: 8},
		Item{key: 22},
		Item{key: 3},
		Item{key: 14},
		Item{key: 22},
	}

	h := heap.New(values, len(values), MaxItem)

	for !h.IsEmpty() {
		i, err := h.Extract()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(i.(Item).key)
	}
	// Output:
	// 22
	// 22
	// 14
	// 8
	// 3
}
