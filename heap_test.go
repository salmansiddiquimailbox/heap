package heap

import (
	"sort"
	"testing"
)

func TestNewHeap(t *testing.T) {
	h := NewHeap()
	if h == nil {
		t.Fatal("NewHeap() returned nil")
	}
	if h.Len() != 0 {
		t.Errorf("Expected empty heap, got size %d", h.Len())
	}
	if !h.Empty() {
		t.Error("Expected heap to be empty")
	}
}

func TestPushAndPop(t *testing.T) {
	h := NewHeap()

	// Test pushing elements
	numbers := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	for _, num := range numbers {
		h.Push(num)
	}

	if h.Len() != len(numbers) {
		t.Errorf("Expected heap size %d, got %d", len(numbers), h.Len())
	}

	// Test that elements come out in ascending order (min-heap)
	expected := make([]int, len(numbers))
	copy(expected, numbers)
	sort.Ints(expected)

	for i, expectedVal := range expected {
		actual := h.Pop()
		if actual != expectedVal {
			t.Errorf("Expected %d at position %d, got %d", expectedVal, i, actual)
		}
	}

	if !h.Empty() {
		t.Error("Expected heap to be empty after popping all elements")
	}
}

func TestPeek(t *testing.T) {
	h := NewHeap()

	// Test peek on empty heap
	if h.Peek() != 0 {
		t.Errorf("Expected 0 for peek on empty heap, got %d", h.Peek())
	}

	// Test peek with elements
	h.Push(5)
	h.Push(2)
	h.Push(8)

	if h.Peek() != 2 {
		t.Errorf("Expected minimum 2, got %d", h.Peek())
	}

	// Peek shouldn't remove elements
	if h.Len() != 3 {
		t.Errorf("Expected heap size 3 after peek, got %d", h.Len())
	}
}

func TestEmptyAndSize(t *testing.T) {
	h := NewHeap()

	if !h.Empty() {
		t.Error("Expected empty heap to be empty")
	}

	if h.Size() != 0 {
		t.Errorf("Expected size 0, got %d", h.Size())
	}

	h.Push(1)

	if h.Empty() {
		t.Error("Expected non-empty heap")
	}

	if h.Size() != 1 {
		t.Errorf("Expected size 1, got %d", h.Size())
	}

	h.Pop()

	if !h.Empty() {
		t.Error("Expected empty heap after pop")
	}

	if h.Size() != 0 {
		t.Errorf("Expected size 0 after pop, got %d", h.Size())
	}
}

func TestHeapProperty(t *testing.T) {
	h := NewHeap()

	// Add elements in random order
	elements := []int{10, 5, 15, 3, 7, 12, 8, 1, 20}
	for _, elem := range elements {
		h.Push(elem)
	}

	// Verify min-heap property: parent should be <= children
	for i := 0; i < len(h.items.items)/2; i++ {
		left := 2*i + 1
		right := 2*i + 2

		if left < len(h.items.items) && h.items.items[i] > h.items.items[left] {
			t.Errorf("Min-heap property violated: parent[%d]=%d > left[%d]=%d",
				i, h.items.items[i], left, h.items.items[left])
		}

		if right < len(h.items.items) && h.items.items[i] > h.items.items[right] {
			t.Errorf("Min-heap property violated: parent[%d]=%d > right[%d]=%d",
				i, h.items.items[i], right, h.items.items[right])
		}
	}
}

func TestLargeHeap(t *testing.T) {
	h := NewHeap()

	// Test with many elements
	for i := 1000; i >= 1; i-- {
		h.Push(i)
	}

	if h.Len() != 1000 {
		t.Errorf("Expected heap size 1000, got %d", h.Len())
	}

	// Verify elements come out in order
	prev := 0
	for !h.Empty() {
		current := h.Pop()
		if current < prev {
			t.Errorf("Heap order violated: %d came after %d", current, prev)
		}
		prev = current
	}
}

func TestRepeatedElements(t *testing.T) {
	h := NewHeap()

	// Test with repeated elements
	elements := []int{3, 1, 3, 1, 2, 2, 3, 1}
	for _, elem := range elements {
		h.Push(elem)
	}

	expected := []int{1, 1, 1, 2, 2, 3, 3, 3}
	for i, expectedVal := range expected {
		actual := h.Pop()
		if actual != expectedVal {
			t.Errorf("Expected %d at position %d, got %d", expectedVal, i, actual)
		}
	}
}

func TestPrintTree(t *testing.T) {
	h := NewHeap()

	// Test PrintTree on empty heap
	h.PrintTree() // Should print "Empty heap"

	// Test PrintTree with elements
	h.Push(5)
	h.Push(2)
	h.Push(8)
	h.Push(1)

	// These should not panic
	h.PrintTree()

	// Verify the heap still works after printing
	if h.Peek() != 1 {
		t.Errorf("Expected minimum 1 after printing, got %d", h.Peek())
	}
}

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap()

	// Test pushing elements
	numbers := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	for _, num := range numbers {
		h.Push(num)
	}

	if h.Len() != len(numbers) {
		t.Errorf("Expected heap size %d, got %d", len(numbers), h.Len())
	}

	// Test that elements come out in descending order (max-heap)
	expected := make([]int, len(numbers))
	copy(expected, numbers)
	sort.Sort(sort.Reverse(sort.IntSlice(expected)))

	for i, expectedVal := range expected {
		actual := h.Pop()
		if actual != expectedVal {
			t.Errorf("Expected %d at position %d, got %d", expectedVal, i, actual)
		}
	}

	if !h.Empty() {
		t.Error("Expected heap to be empty after popping all elements")
	}
}

func TestMaxHeapProperty(t *testing.T) {
	h := NewMaxHeap()

	// Add elements in random order
	elements := []int{10, 5, 15, 3, 7, 12, 8, 1, 20}
	for _, elem := range elements {
		h.Push(elem)
	}

	// Verify max-heap property: parent should be >= children
	for i := 0; i < len(h.items.items)/2; i++ {
		left := 2*i + 1
		right := 2*i + 2

		if left < len(h.items.items) && h.items.items[i] < h.items.items[left] {
			t.Errorf("Max-heap property violated: parent[%d]=%d < left[%d]=%d",
				i, h.items.items[i], left, h.items.items[left])
		}

		if right < len(h.items.items) && h.items.items[i] < h.items.items[right] {
			t.Errorf("Max-heap property violated: parent[%d]=%d < right[%d]=%d",
				i, h.items.items[i], right, h.items.items[right])
		}
	}
}

func TestHeapType(t *testing.T) {
	minHeap := NewHeap()
	maxHeap := NewMaxHeap()

	if minHeap.GetHeapType() != MinHeap {
		t.Error("Expected min-heap type")
	}

	if maxHeap.GetHeapType() != MaxHeap {
		t.Error("Expected max-heap type")
	}
}
