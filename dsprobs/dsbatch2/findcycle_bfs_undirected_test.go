package dsbatch2

import (
	"fmt"
	"testing"
)

type GraphBFSAdjascentTest struct {
	input1, input2 int
}

var graphBFSAdjacentTests = []GraphBFSAdjascentTest{
	{0, 1},
	{1, 2},
	{2, 3},
	{3, 4},
	{4, 0},
}

func TestGraphBFSAdjascent(t *testing.T) {
	bfsUnDirGraph := getBFSUnDirectedGraph()
	visited := make([]int, 10)
	if ok := bfsUnDirGraph.isBFSCycleExists(visited, bfsUnDirGraph.graph.getVertex(0)); !ok {
		t.Error("Cycle not detected")
	} else {
		fmt.Println("Cycle detected")
	}
}

func getBFSUnDirectedGraph() *BFSUnDirectedGraph {
	dfsGraph := InitBFSUnDirectedGraph(5)
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
