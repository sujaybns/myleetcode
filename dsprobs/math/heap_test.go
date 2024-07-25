package math

import (
	"fmt"
	"testing"
)

type HeapTest struct {
	arg1, expected int
}

var heapTests = []HeapTest{
	HeapTest{1, 1},
	HeapTest{10, 2},
	HeapTest{5, 3},
	HeapTest{50, 4},
}

func TestAddToHeap(t *testing.T) {
	heap := Heap{}
	for _, test := range heapTests {
		if heap.AddToHeap(test.arg1); len(heap.items) != test.expected {
			fmt.Println("error")
			//t.Errorf("Output %q not equal to expected %q", len(stack.items), test.expected)
			t.Errorf("Output %d is not as expected %d", len(heap.items), test.expected)
		}
	}
}

func TestRemoveElementFromHeap(t *testing.T) {
	heap := Heap{}
	for _, test := range heapTests {
		if heap.AddToHeap(test.arg1); len(heap.items) != test.expected {
			fmt.Println("error")
			t.Errorf("Output %d is not as expected %d", len(heap.items), test.expected)
		}
	}
	for index := range len(heapTests) {
		initialLength := len(heap.items)
		if heap.RemoveElementFromHeap(); len(heap.items) != initialLength-1 {
			fmt.Printf("error at iteration:%d", index)
			t.Errorf("Output %d is not as expected %d in iteration", len(heap.items), initialLength-1)
		}
	}
}
