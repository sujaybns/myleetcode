package findunion

func initParent(parent *[]int){
	for i :=range len(*parent){
		(*parent)[i] = i
	}
} 

func find(x int,parent *[]int) int{
	if (*parent)[x] == x {
		return x
	} else {
		return find((*parent)[x], parent)
	}
}

func union(x int,y int, parent *[]int ){
	rootX := find(x, parent)
	rootY := find(y, parent)
	if rootX != rootY{
		(*parent)[y] = x
	}
}

/*
// DisjointSet represents a union-find data structure
type DisjointSet struct {
    parent []int
    rank   []int
}

// NewDisjointSet initializes a new DisjointSet with n elements
func NewDisjointSet(n int) *DisjointSet {
    parent := make([]int, n)
    rank := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
        rank[i] = 0
    }
    return &DisjointSet{parent, rank}
}

// Find returns the representative (root) of the set containing element x
func (ds *DisjointSet) Find(x int) int {
    if ds.parent[x] != x {
        ds.parent[x] = ds.Find(ds.parent[x]) // Path compression
    }
    return ds.parent[x]
}

// Union merges the sets containing elements x and y
func (ds *DisjointSet) Union(x, y int) {
    rootX := ds.Find(x)
    rootY := ds.Find(y)
    if rootX != rootY {
        // Union by rank
        if ds.rank[rootX] < ds.rank[rootY] {
            ds.parent[rootX] = rootY
        } else if ds.rank[rootX] > ds.rank[rootY] {
            ds.parent[rootY] = rootX
        } else {
            ds.parent[rootY] = rootX
            ds.rank[rootX]++
        }
    }
}

func main() {
    ds := NewDisjointSet(5)

    ds.Union(0, 2)
    ds.Union(4, 2)
    ds.Union(3, 1)

    fmt.Println(ds.Find(4)) // Output: representative of set containing 4
    fmt.Println(ds.Find(3)) // Output: representative of set containing 3
}
	*/