package leetcode114

import (
	"fmt"
	"testing"
)

type Mytest struct {
	values [3]int
}

var mytests = []*Mytest{
	&Mytest{values: [3]int{1, 2, 3}},
	&Mytest{values: [3]int{1, 3, 2}},
	&Mytest{values: [3]int{2, 1, 3}},
	&Mytest{values: [3]int{2, 3, 1}},
	&Mytest{values: [3]int{3, 1, 2}},
	&Mytest{values: [3]int{3, 2, 1}},
}

func TestPrintInOrder(t *testing.T) {
	fooBar := InitInOrderFooBar(18)

	count := 0
	for _, test := range mytests {
		for i := 0; i < 3; i++ {
			switch val := (*test).values[i]; val {
			case 1:
				count++
				go fooBar.first()
			case 2:
				count++
				go fooBar.second()
			case 3:
				count++
				go fooBar.third()
			}
		}
	}
	fmt.Println("count ", count)
	fooBar.mysync.Wait()
}
