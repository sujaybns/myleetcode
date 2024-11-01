package dsbatch2

import "fmt"

type DFSUnDirectedGraph struct{
	graph *Graph
}

func (g *DFSUnDirectedGraph) isDFSUnDirCycleExists(visited []int, node Vertex, parent Vertex) bool {
	visited[node.val] = 1
	fmt.Println("visited:",visited)
	for _,adjnode := range node.adjascentVertexes{
		if visited[adjnode.val] != 1{
			visited[adjnode.val] = 1
			if g.isDFSUnDirCycleExists(visited,*adjnode, node) == true{
				return true
			}
		} else if adjnode.val != parent.val{
			return true
		}
	}
	return false
}

func (g *DFSUnDirectedGraph) initCycleExists(visited []int, node Vertex, parent Vertex) bool{
	for i:=0;i<len(g.graph.vertices); i++{
		if visited[i] !=0 {
			if (g.isDFSUnDirCycleExists(visited, *g.graph.getVertex(i), *g.graph.getVertex(-1)) == true){
				return true
			}
		}
	}
	return false
} 
 
func InitDFSUnDirectedGraph(n int) *DFSUnDirectedGraph{
	return &DFSUnDirectedGraph{graph: InitGraphAdjascent(n)}
} 