# Heap Wrapper

A simple and intuitive wrapper for Go's `container/heap` package that provides an easy-to-use interface for integer heaps.

## Features

- **Simple API**: Just call `NewHeap()` to create a new min-heap or `NewMaxHeap()` for max-heap
- **Dual Heap Types**: Supports both min-heap and max-heap
- **Type-Safe**: Works specifically with integers
- **Easy Operations**: Push, Pop, Peek, and utility methods
- **Built on container/heap**: Uses Go's standard heap implementation underneath

## Usage

### Basic Operations

```go
package main

import (
    "fmt"
    "github.com/salmansiddiquimailbox/heap"
)

func main() {
    // Create a new min-heap
    minHeap := heap.NewHeap()
    
    // Push integers
    minHeap.Push(5)
    minHeap.Push(2)
    minHeap.Push(8)
    minHeap.Push(1)
    
    // Peek at the minimum (without removing)
    fmt.Println("Minimum:", minHeap.Peek()) // Output: 1
    
    // Pop the minimum
    min := minHeap.Pop()
    fmt.Println("Popped:", min) // Output: 1
    
    // Create a new max-heap
    maxHeap := heap.NewMaxHeap()
    maxHeap.Push(5)
    maxHeap.Push(2)
    maxHeap.Push(8)
    maxHeap.Push(1)
    
    // Peek at the maximum (without removing)
    fmt.Println("Maximum:", maxHeap.Peek()) // Output: 8
    
    // Pop the maximum
    max := maxHeap.Pop()
    fmt.Println("Popped:", max) // Output: 8
    
    // Check heap properties
    fmt.Println("Min-heap size:", minHeap.Size())
    fmt.Println("Max-heap size:", maxHeap.Size())
}
```

### Available Methods

- `NewHeap()` - Creates and returns a new min-heap
- `NewMaxHeap()` - Creates and returns a new max-heap
- `NewHeapWithType(heapType)` - Creates a heap of specified type (MinHeap or MaxHeap)
- `Push(x int)` - Adds an integer to the heap
- `Pop() int` - Removes and returns the top integer (min for min-heap, max for max-heap)
- `Peek() int` - Returns the top integer without removing it
- `Len() int` - Returns the number of items in the heap
- `Size() int` - Alias for Len()
- `Empty() bool` - Returns true if the heap is empty
- `GetHeapType()` - Returns the type of heap (MinHeap or MaxHeap)

## Example

Run the example to see the heap in action:

```bash
go run example/main.go
```

This will demonstrate:
- Creating a heap
- Adding elements
- Peeking at the minimum
- Popping elements in ascending order
- Checking heap properties
- Tree visualization

## Tree Visualization

The heap wrapper includes tree printing functionality to visualize the heap structure:

```go
h := heap.NewHeap()
h.Push(5)
h.Push(2)
h.Push(8)
h.Push(1)

h.PrintArray() // Shows: Heap as array: [1 2 8 5]
h.PrintTree()  // Shows the heap as a tree structure
```

The tree output will look like:
```
Heap as tree:
└── 1
    ┌── 2
    │   └── 5
    └── 8
```

Or in beautiful hierarchical format:
```
Heap as hierarchical tree:
1
├── 2
│   └── 5
└── 8
```

This makes it easy to understand the heap structure and verify that the heap property is maintained.

## Implementation Details

The wrapper implements the `container/heap.Interface` and provides a min-heap of integers. The underlying implementation uses Go's standard `container/heap` package, ensuring reliability and performance.

## Requirements

- Go 1.11 or later
- No external dependencies (uses only standard library) 