package dsbatch2

// Graph - set of pointer to vertices
type Graph struct {
	vertices []*Vertex
}

// Vertex - vertex value and its adjascent nodes
type Vertex struct {
	val               int
	adjascentVertexes []*Vertex
}

// Add vertex - add a vertex to the Graph
func (g *Graph) addVertex(val int) {
	g.vertices = append(g.vertices, &Vertex{val: val})
}

// Add edge - add a vertex(To) to the adjascency list of vertex (From)
// check if the vertex is in Graph
// check if the edge already exists
func (g *Graph) addEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if g.isValidVertices(from, to) && g.isNonDuplicateVertex(fromVertex, to) {
		fromVertex.adjascentVertexes = append(fromVertex.adjascentVertexes, toVertex)
	}
}

func (g *Graph) isValidVertices(from, to int) bool {
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

func (g *Graph) isNonDuplicateVertex(fromVertex *Vertex, to int) bool {
	for _, vertex := range fromVertex.adjascentVertexes {
		if vertex.val == to {
			return false
		}
	}
	return true
}

func (g *Graph) getVertex(val int) *Vertex {
	if val == -1 {
		return &Vertex{val:-1}
	}
	for _, vertex := range g.vertices {
		if vertex.val == val {
			return vertex
		}
	}
	return nil
}

func InitGraphAdjascent(numberOfvertices int) *Graph {
	graph := &Graph{}
	for i := range numberOfvertices {
		graph.addVertex(i)
	}
	return graph
}
