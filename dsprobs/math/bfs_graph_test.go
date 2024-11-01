package math

import (
	"fmt"
	"slices"
	"testing"
)

type BFSGraphAdjascentTest struct {
	input1, input2 int
}

var bfsGraphAdjacentTests = []BFSGraphAdjascentTest{
	{0, 1},
	{1, 2},
	{2, 3},
	{3, 4},
	{4, 0},
}

type BFSGraphTestResult struct {
	output []int
}

var bfsGraphTestResult = BFSTreeTestResult{output: []int{0, 1, 2, 3, 4}}

func TestBFSGraph(t *testing.T) {

	graph := initBFSGraphAdjascent(5)
	for _, test := range graphAdjacentTests {
		if graph.addEdge(test.input1, test.input2); len(graph.vertices[test.input1].adjascentVertexes) < 1 {
			t.Error("Actual and expected values do not match")
		}
	}
	fmt.Println("\nBFS traversal")
	if outputList := graph.BfsGraphTraversal(graph.getVertex(0)); slices.Equal(bfsGraphTestResult.output, BFSGraphConvertSlice[int](outputList)) != true {
		t.Errorf("Expected and actual values are different")
	}
}

func BFSGraphConvertSlice[E any](in []any) (out []E) {
	out = make([]E, 0, len(in))
	for _, v := range in {
		out = append(out, v.(E))
	}
	return
}
