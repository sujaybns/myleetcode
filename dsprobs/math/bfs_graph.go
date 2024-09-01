package math

// import "fmt"

type BFSGraph struct {
	vertices []*BFSVertex
}

// Vertex - vertex value and its adjascent nodes
type BFSVertex struct {
	val               int
	adjascentVertexes []*BFSVertex
}

// Add vertex - add a vertex to the Graph
func (g *BFSGraph) addVertex(val int) {
	g.vertices = append(g.vertices, &BFSVertex{val: val})
}

// Add edge - add a vertex(To) to the adjascency list of vertex (From)
// check if the vertex is in Graph
// check if the edge already exists
func (g *BFSGraph) addEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if g.isValidVertices(from, to) && g.isNonDuplicateVertex(fromVertex, to) {
		fromVertex.adjascentVertexes = append(fromVertex.adjascentVertexes, toVertex)
	}
}

func (g *BFSGraph) isValidVertices(from, to int) bool {
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

func (g *BFSGraph) isNonDuplicateVertex(fromVertex *BFSVertex, to int) bool {
	for _, vertex := range fromVertex.adjascentVertexes {
		if vertex.val == to {
			return false
		}
	}
	return true
}

func (g *BFSGraph) getVertex(val int) *BFSVertex {
	for _, vertex := range g.vertices {
		if vertex.val == val {
			return vertex
		}
	}
	return nil
}

func (g *BFSGraph) BfsGraphTraversal(startVertex *BFSVertex)  []interface{}{
	var newVistedList = make([]int,len(g.vertices))
	newVistedList[startVertex.val] = 1
	queue := Queue{} 
	queue.Enqueue(startVertex)
	// newVistedList[startVertex.val] = 1
	var newList []interface{}
	for !queue.IsEmpty(){
		size := queue.Size()
		for i := 0;i < size ; i++{
			item := queue.Dequeue().(*BFSVertex)
			newList = append(newList, item.val)
			for _,v := range item.adjascentVertexes{
				if newVistedList[v.val] != 1{
					newVistedList[v.val] = 1
					queue.Enqueue(v)
				}
			}
		}
	}
	return newList
}


func initBFSGraphAdjascent(numberOfvertices int) *BFSGraph {
	graph := &BFSGraph{}
	for i := range numberOfvertices {
		graph.addVertex(i)
	}
	return graph
}