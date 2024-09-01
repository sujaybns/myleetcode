package math

import (
	"fmt"
	"slices"
	"testing"
)

type DFSGraphAdjascentTest struct {
	input1, input2 int
}

var dfsGraphAdjacentTests = []DFSGraphAdjascentTest{
	{0, 1},
	{1, 2},
	{2, 3},
	{3, 4},
	{4, 0},
}

type DFSGraphTestResult struct {
	output []int
}

var dfsGraphTestResult = DFSGraphTestResult{output: []int{0,1,2,3,4}}

func TestDFSGraph(t *testing.T) {

	graph := initDFSGraphAdjascent(5)
	for _, test := range graphAdjacentTests {
		if graph.addEdge(test.input1, test.input2); len(graph.vertices[test.input1].adjascentVertexes) < 1 {
			t.Error("Actual and expected values do not match")
		}
	}
	fmt.Println("\nBFS traversal")
	visitedList := [5]int{1,0,0,0,0} 
	list := []int{}
	if graph.DfsGraphTraversal(graph.getVertex(0),visitedList,&list); slices.Equal(bfsGraphTestResult.output, list) != true {
		t.Errorf("Expected and actual values are different")
	}
}