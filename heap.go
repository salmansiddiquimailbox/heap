package heap

import (
	"fmt"

	"container/heap"
)

// HeapType represents the type of heap (min or max)
type HeapType int

const (
	MinHeap HeapType = iota
	MaxHeap
)

// intSlice implements container/heap.Interface
type intSlice struct {
	items    []int
	heapType HeapType
}

func (h intSlice) Len() int { return len(h.items) }

func (h intSlice) Less(i, j int) bool {
	if h.heapType == MinHeap {
		return h.items[i] < h.items[j]
	}
	return h.items[i] > h.items[j] // MaxHeap
}

func (h intSlice) Swap(i, j int) { h.items[i], h.items[j] = h.items[j], h.items[i] }

func (h *intSlice) Push(x interface{}) {
	h.items = append(h.items, x.(int))
}

func (h *intSlice) Pop() interface{} {
	old := h.items
	n := len(old)
	x := old[n-1]
	old[n-1] = 0 // avoid memory leak
	h.items = old[0 : n-1]
	return x
}

// IntHeap represents a heap of integers (min or max)
type IntHeap struct {
	items *intSlice
}

// NewHeap creates and returns a new min-heap of integers
func NewHeap() *IntHeap {
	return NewHeapWithType(MinHeap)
}

// NewMaxHeap creates and returns a new max-heap of integers
func NewMaxHeap() *IntHeap {
	return NewHeapWithType(MaxHeap)
}

// NewHeapWithType creates and returns a new heap of the specified type
func NewHeapWithType(heapType HeapType) *IntHeap {
	h := &IntHeap{
		items: &intSlice{
			items:    make([]int, 0),
			heapType: heapType,
		},
	}
	heap.Init(h.items)
	return h
}

// Push adds an integer to the heap
func (h *IntHeap) Push(x int) {
	heap.Push(h.items, x)
}

// Pop removes and returns the top integer from the heap (min for min-heap, max for max-heap)
func (h *IntHeap) Pop() int {
	return heap.Pop(h.items).(int)
}

// Peek returns the top integer without removing it (returns 0 if heap is empty)
func (h *IntHeap) Peek() int {
	if h.Len() == 0 {
		return 0
	}
	return h.items.items[0]
}

// Len returns the number of items in the heap
func (h *IntHeap) Len() int {
	return len(h.items.items)
}

// Empty returns true if the heap is empty
func (h *IntHeap) Empty() bool {
	return h.Len() == 0
}

// Size returns the number of items in the heap (alias for Len)
func (h *IntHeap) Size() int {
	return h.Len()
}

// GetHeapType returns the type of the heap (MinHeap or MaxHeap)
func (h *IntHeap) GetHeapType() HeapType {
	return h.items.heapType
}

// PrintTree prints the heap as a hierarchical tree structure
func (h *IntHeap) PrintTree() {
	if h.Empty() {
		fmt.Println("Empty heap")
		return
	}

	heapType := "min-heap"
	if h.GetHeapType() == MaxHeap {
		heapType = "max-heap"
	}
	fmt.Printf("Heap as tree (%s):\n", heapType)
	h.printTreeHelper(0, "", true)
}

// printTreeHelper is a helper function to recursively print the tree
func (h *IntHeap) printTreeHelper(index int, prefix string, isLeft bool) {
	if index >= len(h.items.items) {
		return
	}

	// Print current node
	fmt.Printf("%s", prefix)
	if isLeft {
		fmt.Print("└── ")
	} else {
		fmt.Print("┌── ")
	}
	fmt.Printf("%d\n", h.items.items[index])

	// Calculate new prefix for children
	newPrefix := prefix
	if isLeft {
		newPrefix += "    "
	} else {
		newPrefix += "│   "
	}

	// Check if we have children
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	hasLeft := leftChild < len(h.items.items)
	hasRight := rightChild < len(h.items.items)

	// Print left child first
	if hasLeft {
		h.printTreeHelper(leftChild, newPrefix, !hasRight)
	}

	// Print right child
	if hasRight {
		h.printTreeHelper(rightChild, newPrefix, true)
	}
}

// PrintArray prints the heap as a simple array
func (h *IntHeap) PrintArray() {
	if h.Empty() {
		fmt.Println("Empty heap")
		return
	}

	heapType := "min-heap"
	if h.GetHeapType() == MaxHeap {
		heapType = "max-heap"
	}
	fmt.Printf("Heap as array (%s): %v\n", heapType, h.items.items)
}
