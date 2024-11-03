package mydijkstra

import "fmt"

type PriorityQueue []Node

func (pq PriorityQueue) Len() int { return len(pq) }

// For a min-heap, please use pq[i].DistanceFromSource < pq[j].DistanceFromSource.
// To create a max-heap, please use pq[i].DistanceFromSource > pq[j].DistanceFromSource.
// Here Priority can be a weighed value like DistanceFromSource
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].DistanceFromSource < pq[j].DistanceFromSource
}

// Print() prints the item with the highest priority (min-heap).
func (pq PriorityQueue) Print() {
	for i := 0; i < len(pq); i++ {
		fmt.Println(pq[i].Value)
		fmt.Println(pq[i].DistanceFromSource)
	}
}

// Swap() swaps elements with indices i and j.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push() adds an item to priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Node)
	*pq = append(*pq, item)
}

// Pop() removes the item with highest priority (min-heap).
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
