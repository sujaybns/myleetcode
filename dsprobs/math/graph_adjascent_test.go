package math

import (
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

func TestGraphAdjascent(t *testing.T) {
	graph := initGraphAdjascent(5)
	for _, test := range graphAdjacentTests {
		if graph.addEdge(test.input1, test.input2); len(graph.vertices[test.input1].adjascentVertexes) < 1 {
			t.Error("Actual and expected values do not match")
		}
	}
	if graph.addVertex(5); len(graph.vertices) != 6 {
		t.Error("Actual and expected values do not match")
	}

}
