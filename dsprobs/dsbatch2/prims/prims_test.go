package prims

import (
	"fmt"
"testing"
)

func TestPrims(t *testing.T){
	graph := initGraph()
	node := initNode(0)
	node.addAdjNodeWithDistance(1,2)
	node.addAdjNodeWithDistance(2,1)
	graph.addNode(node)

	node1 := initNode(1)
	node1.addAdjNodeWithDistance(0,2)
	node1.addAdjNodeWithDistance(3,3)
	graph.addNode(node1)

	node2 := initNode(2)
	node2.addAdjNodeWithDistance(0,1)
	node2.addAdjNodeWithDistance(3,4)
	graph.addNode(node2)

	node3 := initNode(3)
	node3.addAdjNodeWithDistance(2,4)
	node3.addAdjNodeWithDistance(1,3)
	graph.addNode(node3)

	value := graph.prims()
	expected := 6
	if value != expected {
		t.Errorf("prims() = %d; want %d", value, expected)
	}
	fmt.Printf("Value : %v", value)
}
