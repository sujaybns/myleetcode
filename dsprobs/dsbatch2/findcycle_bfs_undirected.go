package dsbatch2

import (
	"fmt"

	"sujay.com/challenges/math"
)

type BFSUnDirectedGraph struct {
	graph *Graph
}

type Pair struct {
	v1     *Vertex
	parent *Vertex
}

func (g *BFSUnDirectedGraph) isBFSCycleExists(visited []int, node *Vertex) bool {
	q := math.Queue{}
	fmt.Println("q:", q)
	q.Enqueue(&Pair{v1: node, parent: g.graph.getVertex(-1)})
	q.IsEmpty()
	visited[node.val] = 1
	for !q.IsEmpty() {
		size := q.Size()
		for i := 0; i < size; i++ {
			pair := q.Dequeue().(*Pair)
			newNode := pair.v1
			for _, adjnode := range newNode.adjascentVertexes {
				if visited[adjnode.val] != 1 {
					visited[adjnode.val] = 1
					q.Enqueue(&Pair{v1: adjnode, parent: newNode})
				} else if adjnode.val != newNode.val {
					return true
				}
			}
		}
	}
	return false
}


func InitBFSUnDirectedGraph(n int) *BFSUnDirectedGraph {
	return &BFSUnDirectedGraph{graph: InitGraphAdjascent(n)}
}
