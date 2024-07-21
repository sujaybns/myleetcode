package math

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
	fmt.Println("Enqueued")
	fmt.Println(q.items)
}

func (q *Queue) Dequeue() int {
	item := q.items[0]
	q.items = q.items[1:]
	fmt.Println("Dequeued")
	fmt.Println(q.items)
	return item
}
