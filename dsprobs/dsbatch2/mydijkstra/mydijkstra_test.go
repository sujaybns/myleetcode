package mydijkstra

import (
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := &Graph{nodes: []*Node{}}

	nodeA := InitNode("A", 0, 2)
	nodeA.addAdjNodeWithDistance("B", 4)
	nodeA.addAdjNodeWithDistance("C", 2)

	nodeB := InitNode("B", 10000, 3)
	nodeB.addAdjNodeWithDistance("C", 3)
	nodeB.addAdjNodeWithDistance("D", 2)
	nodeB.addAdjNodeWithDistance("E", 3)

	nodeC := InitNode("C", 10000, 3)
	nodeC.addAdjNodeWithDistance("B", 1)
	nodeC.addAdjNodeWithDistance("D", 4)
	nodeC.addAdjNodeWithDistance("E", 5)

	nodeD := InitNode("D", 10000, 0)

	nodeE := InitNode("E", 10000, 1)
	nodeE.addAdjNodeWithDistance("D", 1)

	graph.addNode(nodeA)
	graph.addNode(nodeB)
	graph.addNode(nodeC)
	graph.addNode(nodeD)
	graph.addNode(nodeE)

	fmt.Println("graph", graph)

	settled := make(map[string]int, 5)
	graph.dijkstraAlgorithm(settled, "A")
}