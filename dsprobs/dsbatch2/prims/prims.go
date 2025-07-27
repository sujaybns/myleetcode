package prims

import "container/heap"

type Node struct {
	val            int
	adjNodeWithVal map[int]int
}

type Edge struct {
	distance int
	value    int
	parent   int
}

type Graph struct {
	nodes []*Node
}

func (g *Graph) getNode(node int) *Node {
	for _, gNode := range g.nodes {
		if gNode.val == node {
			return gNode
		}
	}
	return nil
}

func (g *Graph) addNode(node *Node) {
	g.nodes = append(g.nodes, node)
}

func (node *Node) addAdjNodeWithDistance(nodeval int, distance int) {
	node.adjNodeWithVal[nodeval] = distance
}

func initNode(value int) *Node {
	return &Node{val: value, adjNodeWithVal: make(map[int]int)}
}

func initGraph() *Graph {
	graph := &Graph{nodes: []*Node{}}
	return graph
}

func (graph *Graph) prims() int {

	visited := make(map[int]int)
	for _, node := range graph.nodes {
		visited[node.val] = 0
	}
	// heap
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Edge{0, 0, -1})
	sum :=0
	for pq.Len() != 0 {
		edge := heap.Pop(pq).(Edge)
		if visited[edge.value] == 1 {
			continue
		}
		visited[edge.value] = 1
		sum = sum + edge.distance
		for adjNode, distance := range graph.getNode(edge.value).adjNodeWithVal {
			if visited[adjNode] == 0{
				heap.Push(pq,Edge{distance: distance,value: adjNode,parent: edge.value})
			} 

		}
	}
	return sum
}
