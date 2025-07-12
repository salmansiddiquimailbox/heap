package main

import (
	"fmt"

	"github.com/siddiqui-io/heap"
)

func main() {
	// Create a new heap
	h := heap.NewHeap()

	fmt.Println("=== Integer Heap Example ===")

	// Push some integers
	numbers := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	fmt.Printf("Adding numbers: %v\n", numbers)

	for _, num := range numbers {
		h.Push(num)
		fmt.Printf("Pushed %d, heap size: %d, peek: %d\n", num, h.Size(), h.Peek())
	}

	fmt.Println("\n=== Heap Operations ===")

	// Show heap properties
	fmt.Printf("Heap size: %d\n", h.Size())
	fmt.Printf("Is empty: %t\n", h.Empty())
	fmt.Printf("Peek (minimum): %d\n", h.Peek())

	// Pop all elements (they will come out in ascending order)
	fmt.Println("\nPopping all elements (min-heap order):")
	for !h.Empty() {
		min := h.Pop()
		fmt.Printf("Popped: %d, remaining size: %d\n", min, h.Size())
	}

	fmt.Printf("Final heap size: %d\n", h.Size())
	fmt.Printf("Is empty: %t\n", h.Empty())

	// Demonstrate with a new heap
	fmt.Println("\n=== Another Example ===")
	h2 := heap.NewHeap()

	// Add some numbers in different order
	h2.Push(10)
	h2.Push(5)
	h2.Push(15)
	h2.Push(3)
	h2.Push(7)

	fmt.Printf("Heap size: %d\n", h2.Size())
	fmt.Printf("Minimum element: %d\n", h2.Peek())

	// Pop a few elements
	fmt.Println("Popping first 3 elements:")
	for i := 0; i < 3; i++ {
		if !h2.Empty() {
			fmt.Printf("Popped: %d\n", h2.Pop())
		}
	}

	fmt.Printf("Remaining heap size: %d\n", h2.Size())
	if !h2.Empty() {
		fmt.Printf("New minimum: %d\n", h2.Peek())
	}
}
