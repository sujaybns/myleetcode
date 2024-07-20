package math

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
	fmt.Println("pushed")
}

func (s *Stack) Pop() int {
	l := len(s.items) - 1
	item := s.items[l]
	s.items = s.items[:l]
	fmt.Println("popped")
	return item
}
