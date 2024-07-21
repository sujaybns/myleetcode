package math

import "testing"

type HeapTest struct {
	arg1, output int
}

var heapTests = [] HeapTest {
	HeapTest{1,1},
	HeapTest{10,2},
	HeapTest{5,3},
	HeapTest{50,4},
}

func TestAddToHeap( t *testing.T){
	heap := Heap{}
	for _, test := range(heapTests){
		heap.AddToHeap(test.arg1)
	}
}
