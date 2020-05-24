# heap
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/fsmiamoto/heap)
[![Go Report Card](https://goreportcard.com/badge/github.com/fsmiamoto/heap)](https://goreportcard.com/report/github.com/fsmiamoto/heap)
[![GoCover](https://gocover.io/_badge/github.com/fsmiamoto/heap)](https://gocover.io/github.com/fsmiamoto/heap)
![Test](https://github.com/fsmiamoto/heap/workflows/Test/badge.svg?branch=master)

A generic binary heap data structure in Go.

It works with *any* type as long as you provide a comparison function.

The package defines two example functions for building a Max or Min Heap of integers.

## Examples

- Heapsort
```go 
func main() {
    values := []interface{}{40, 30, 50, 100, 15}

    // You can specify an initial capacity for the heap,
    // len(values) in this case, which helps to avoid reallocations.
    // Also, this heap uses the package defined MinInt comparison function,
    // that builds a MinHeap of integers
    h := heap.New(values, len(values), heap.MinInt)

    var sorted []int
    for !h.IsEmpty() {
        v, err := h.Extract()
        if err != nil {
            log.Fatal(err)
        }
        sorted = append(sorted, v.(int))
    }
    fmt.Println(sorted) // [15 30 40 50 100]
}
```
