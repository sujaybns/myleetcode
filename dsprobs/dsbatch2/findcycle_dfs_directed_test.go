package dsbatch2

import (
	"fmt"
	"testing"
)

type GraphAdjascentTest struct {
	input1, input2 int
}

var graphAdjacentTests = []GraphAdjascentTest{
	{0, 1},
	{1, 2},
	{2, 3},
	{3, 4},
	{4, 0},
}

func TestGraphDFSDirectedAdjascent(t *testing.T) {
	dfsGraph := getGraph()
	visited := make([]int, 10)
	path := make([]int, 10)
	if ok := dfsGraph.initCycleExists(visited, path); !ok {
		t.Error("Cycle not detected")
	} else {
		fmt.Println("Cycle detected")
	}
}

func getGraph() *DFSDirectedGraph {
	dfsGraph := InitDFSDirectedGraph(5)
	for _, test := range graphAdjacentTests {
		if dfsGraph.graph.addEdge(test.input1, test.input2); len(dfsGraph.graph.vertices[test.input1].adjascentVertexes) < 1 {
			fmt.Println("Actual and expected values do not match")
		}
	}
	if dfsGraph.graph.addVertex(5); len(dfsGraph.graph.vertices) != 6 {
		fmt.Println("Actual and expected values do not match")
	}
	return dfsGraph
}
