package findunion

import (
	"sort"
	"fmt"
)

type Edge struct{
	start, end int
	distance int
}

func kruskals(edges *[]Edge) int {
	sort.Slice(*edges,func(i,j int) bool{
		return (*edges)[i].distance < (*edges)[j].distance
	})
	fmt.Printf("Length of edges:%v", len(*edges))
	parent :=make([]int, len(*edges)) 
	initParent(&parent)
	minCost := 0
	for _,edge := range (*edges) {
		u := (edge).start
		v := (edge).end
		if find(u,&parent) == find(v,&parent) { 
			continue
		}
		union(u,v, &parent)
		minCost += edge.distance
	}
	fmt.Printf("Min cost is %v", minCost)
	return minCost
}