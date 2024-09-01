package math

// import "fmt"

type DFSGraph struct {
	vertices []*DFSVertex
}

// Vertex - vertex value and its adjascent nodes
type DFSVertex struct {
	val               int
	adjascentVertexes []*DFSVertex
}

// Add vertex - add a vertex to the Graph
func (g *DFSGraph) addVertex(val int) {
	g.vertices = append(g.vertices, &DFSVertex{val: val})
}

// Add edge - add a vertex(To) to the adjascency list of vertex (From)
// check if the vertex is in Graph
// check if the edge already exists
func (g *DFSGraph) addEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if g.isValidVertices(from, to) && g.isNonDuplicateVertex(fromVertex, to) {
		fromVertex.adjascentVertexes = append(fromVertex.adjascentVertexes, toVertex)
	}
}

func (g *DFSGraph) isValidVertices(from, to int) bool {
	isFromValid := false
	isToValid := false

	for _, vertex := range g.vertices {
		if vertex.val == from {
			isFromValid = true
		}
		if vertex.val == to {
			isToValid = true
		}
	}
	return isFromValid && isToValid
}

func (g *DFSGraph) isNonDuplicateVertex(fromVertex *DFSVertex, to int) bool {
	for _, vertex := range fromVertex.adjascentVertexes {
		if vertex.val == to {
			return false
		}
	}
	return true
}

func (g *DFSGraph) getVertex(val int) *DFSVertex {
	for _, vertex := range g.vertices {
		if vertex.val == val {
			return vertex
		}
	}
	return nil
}

func (g *DFSGraph) DfsGraphTraversal(root *DFSVertex, newVistedList [5]int, dfslist *[]int) {
	if root == nil {
		return
	} 
	// else {
		*dfslist = append(*dfslist, root.val)
		for _,v := range root.adjascentVertexes{
			if newVistedList[v.val] != 1 {
				newVistedList[v.val] = 1
				g.DfsGraphTraversal(v,newVistedList,dfslist)
			}
		}
	// }
}



func initDFSGraphAdjascent(numberOfvertices int) *DFSGraph {
	graph := &DFSGraph{}
	for i := range numberOfvertices {
		graph.addVertex(i)
	}
	return graph
}