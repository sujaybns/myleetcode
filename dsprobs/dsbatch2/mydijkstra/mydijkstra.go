package mydijkstra

import (
	"container/heap"
	"fmt"
)

type Graph struct {
	nodes []*Node
}

func (g *Graph) addNode(node *Node) {
	g.nodes = append(g.nodes, node)
}

func (g *Graph) getNode(node string) *Node {
	for _, gNode := range g.nodes {
		if gNode.Value == node {
			return gNode
		}
	}
	return nil
}

type Node struct {
	Value                string
	DistanceFromSource   int
	AdjNodesWithDistance map[string]int
}

func (node *Node) String() {
	fmt.Println("value:", node.Value)
	fmt.Println("distanceFromSource:", node.DistanceFromSource)
	fmt.Printf("adjNodesWithDistance:%+v", node.AdjNodesWithDistance)
}

func (node *Node) addAdjNodeWithDistance(adjNode string, adjNodedistance int) {
	node.AdjNodesWithDistance[adjNode] = adjNodedistance
}

func InitNode(newValue string, newDistance int, numberOfAdjNodes int) *Node {
	return &Node{Value: newValue, DistanceFromSource: newDistance, AdjNodesWithDistance: make(map[string]int, numberOfAdjNodes)}
}

func (g Graph) dijkstraAlgorithm(settled map[string]int, start string) {
	distances := make(map[string]int)
	for _, name := range g.nodes {
		distances[name.Value] = 100000
	}
	distances[start] = 0

	visited := make(map[string]int)
	for _, name := range g.nodes {
		visited[name.Value] = 0
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Node{Value: start, DistanceFromSource: 0})
	for pq.Len() != 0 {
		node := heap.Pop(pq).(Node)
		fmt.Println("node(val):", node.Value)
		fmt.Println("node(dist):", node.DistanceFromSource)
		if visited[node.Value] == 1 {
			continue
		}
		visited[node.Value] = 1
		for adjNode, distance := range g.getNode(node.Value).AdjNodesWithDistance {
			fmt.Println("adjnode(start):", adjNode)
			fmt.Println("distance(start):", distance)
			if distances[adjNode] > distances[node.Value]+distance {
				distances[adjNode] = distances[node.Value] + distance
				fmt.Println("adjnode(end):", adjNode)
				fmt.Println("distance(end):", distances[adjNode])
				heap.Push(pq, Node{Value: adjNode, DistanceFromSource: distances[adjNode]})
			}
		}
		pq.Print()
	}
	fmt.Println("distances:", distances)
}