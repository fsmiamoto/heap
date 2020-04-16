package main

import (
	"fmt"

	"github.com/fsmiamoto/heap"
)

func main() {
	values := []interface{}{40, 30, 50, 100, 15}

	h := heap.New(values, len(values), heap.MinInt)

	var sorted []int

	for {
		v, err := h.Extract()
		if err != nil {
			break
		}
		sorted = append(sorted, v.(int))
	}

	fmt.Println(sorted)
}
