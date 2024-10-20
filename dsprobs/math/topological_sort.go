package math

type TopologicalGraph struct {
	vertices []*TopologicalVertex
}

// Vertex - vertex value and its adjascent nodes
type TopologicalVertex struct {
	val               int
	adjascentVertexes []*TopologicalVertex
}

// Add vertex - add a vertex to the Graph
func (g *TopologicalGraph) addVertex(val int) {
	g.vertices = append(g.vertices, &TopologicalVertex{val: val})
}

// Add edge - add a vertex(To) to the adjascency list of vertex (From)
// check if the vertex is in Graph
// check if the edge already exists
func (g *TopologicalGraph) addEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if g.isValidVertices(from, to) && g.isNonDuplicateVertex(fromVertex, to) {
		fromVertex.adjascentVertexes = append(fromVertex.adjascentVertexes, toVertex)
	}
}

func (g *TopologicalGraph) isValidVertices(from, to int) bool {
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

func (g *TopologicalGraph) isNonDuplicateVertex(fromVertex *TopologicalVertex, to int) bool {
	for _, vertex := range fromVertex.adjascentVertexes {
		if vertex.val == to {
			return false
		}
	}
	return true
}

func (g *TopologicalGraph) getVertex(val int) *TopologicalVertex {
	for _, vertex := range g.vertices {
		if vertex.val == val {
			return vertex
		}
	}
	return nil
}

func (g *TopologicalGraph) topoSort() []int {
	visited := make([]int, len(g.vertices))
	stack := &Stack{}
	for i := 0; i < len(g.vertices); i++ {
		if visited[i] == 0 {
			g.dfsTopologicalSort(i, visited, stack)
		}
	}

	size := len(stack.items)
	result := make([]int, 0)
	for i := 0; i < size; i++ {
		result = append(result, stack.Pop())
	}
	return result
}

func (g *TopologicalGraph) dfsTopologicalSort(node int, visited []int, stack *Stack) {
	visited[node] = 1
	for _, adjNode := range g.vertices[node].adjascentVertexes {
		if visited[adjNode.val] == 0 {
			g.dfsTopologicalSort(adjNode.val, visited, stack)
		}
	}
	stack.Push(node)
}

func initTopologicalGraphAdjascent(numberOfvertices int) *TopologicalGraph {
	graph := &TopologicalGraph{}
	for i := range numberOfvertices {
		graph.addVertex(i)
	}
	return graph
}
