package math

import (
	"fmt"
	"slices"
	"testing"
)

type TopologicalGraphAdjascentTest struct {
	input1, input2 int
}

var topologicalGraphAdjacentTests = []TopologicalGraphAdjascentTest{
	{0, 1},
	{1, 2},
	{2, 3},
	{3, 4},
	{4, 0},
}

type TopologicalGraphTestResult struct {
	output []int
}

var topologicalGraphTestResult = TopologicalGraphTestResult{output: []int{0, 1, 2, 3, 4}}

func TestTopologicalGraph(t *testing.T) {
	graph := initTopologicalGraphAdjascent(5)
	for _, test := range topologicalGraphAdjacentTests {
		if graph.addEdge(test.input1, test.input2); len(graph.vertices[test.input1].adjascentVertexes) < 1 {
			t.Error("Actual and expected values do not match")
		}
	}
	fmt.Println("\nTopological traversal")
	// list := make([]int, 0)
	list := []int{}
	if list = graph.topoSort(); slices.Equal(topologicalGraphTestResult.output, list) != true {
		t.Errorf("Expected and actual values are different")
	}
	fmt.Println("list", list)

}
