package math

// import "fmt"

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
	// fmt.Println("Enqueued")
	// fmt.Println(q.items)
}

func (q *Queue) Dequeue() interface{} {
	item := q.items[0]
	q.items = q.items[1:]
	// fmt.Println("Dequeued")
	// fmt.Println(q.items)
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}
