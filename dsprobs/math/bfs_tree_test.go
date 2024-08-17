package math

import (
	"fmt"
	"slices"
	"testing"
)

type BFSTreeTest struct {
	arg1   int
	output int
}

type BFSTreeTestResult struct {
	output []int
}

var bfsTreeTests = []BFSTreeTest{
	{100, 0},
	{80, 0},
	{110, 0},
	{70, 0},
}
var bfsTreeTestResult = BFSTreeTestResult{output: []int{100, 80, 110, 70}}

func TestBFSTree(t *testing.T) {

	bfsTre := initBFSTree()
	for _, test := range bfsTreeTests {
		if bfsTre.Insert(test.arg1); bfsTre.Search(test.arg1) != true {
			t.Errorf("Expected and actual values are different")
		}
	}
	fmt.Println("\nBFS traversal")
	if outputList := bfsTre.BFSTraversal(); slices.Equal(bfsTreeTestResult.output, ConvertSlice[int](outputList)) != true {
		t.Errorf("Expected and actual values are different")
	}
}

func ConvertSlice[E any](in []any) (out []E) {
	out = make([]E, 0, len(in))
	for _, v := range in {
		out = append(out, v.(E))
	}
	return
}
