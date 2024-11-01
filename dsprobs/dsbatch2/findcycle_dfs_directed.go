package dsbatch2

import "fmt"

type DFSDirectedGraph struct {
	graph *Graph
}

func (g *DFSDirectedGraph) isDFSDirCycleExists(visited []int, path []int, node Vertex) bool {
	visited[node.val] = 1
	path[node.val] = 1
	fmt.Println("visited:", visited)
	fmt.Println("path:", path)
	for _, adjnode := range node.adjascentVertexes {
		if visited[adjnode.val] != 1 {
			visited[adjnode.val] = 1
			if g.isDFSDirCycleExists(visited, path, *adjnode) == true {
				return true
			}
		} else if path[adjnode.val] == 1 {
			return true
		}
	}
	path[node.val] = 0
	return false
}

func (g *DFSDirectedGraph) initCycleExists(visited []int, path []int) bool {
	for i := 0; i < len(g.graph.vertices); i++ {
		if visited[i] != 1 {
			if g.isDFSDirCycleExists(visited, path, *g.graph.getVertex(i)) == true {
				return true
			}
		}
	}
	return false
}

func InitDFSDirectedGraph(n int) *DFSDirectedGraph {
	return &DFSDirectedGraph{graph: InitGraphAdjascent(n)}
}
