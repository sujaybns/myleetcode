package dsbatch2

import (
	"fmt"
	"testing"
)

type GraphDFSUDAdjascentTest struct {
	input1, input2 int
}

var graphDFSUDAdjacentTests = []GraphDFSUDAdjascentTest{
	{0, 1},
	{1, 2},
	{2, 3},
	{3, 4},
	{4, 0},
}

func TestDFSUndirectedGraphAdjascent(t *testing.T) {
	dfsGraph := getDFSUnDirectedGraph(t)
	visited := make([]int, 10)
	result := dfsGraph.initCycleExists(visited, *dfsGraph.graph.getVertex(0), *dfsGraph.graph.getVertex(-1))
	fmt.Println("result:", result)
}

func getDFSUnDirectedGraph(t *testing.T) *DFSUnDirectedGraph {
	dfsGraph := InitDFSUnDirectedGraph(5)
	for _, test := range graphAdjacentTests {
		if dfsGraph.graph.addEdge(test.input1, test.input2); len(dfsGraph.graph.vertices[test.input1].adjascentVertexes) < 1 {
			t.Error("Actual and expected values do not match")
		}
	}
	if dfsGraph.graph.addVertex(5); len(dfsGraph.graph.vertices) != 6 {
		t.Error("Actual and expected values do not match")
	}
	return dfsGraph
}
