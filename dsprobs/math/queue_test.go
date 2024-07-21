package math

import (
	"fmt"
	"testing"
)

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type queueTest struct {
	arg1, expected int
}

var queueTests = []queueTest{
	queueTest{2, 1},
	queueTest{4, 2},
	queueTest{6, 3},
	queueTest{3, 4},
}

func TestEnqueue(t *testing.T) {
	myQueue := Queue{}
	fmt.Println(myQueue)
	for _, test := range queueTests {
		fmt.Println(test.arg1)
		if myQueue.Enqueue(test.arg1); len(myQueue.items) != test.expected {
			fmt.Println("error")
			//t.Errorf("Output %q not equal to expected %q", len(stack.items), test.expected)
			t.Errorf("Output %d is not as expected %d", len(myQueue.items), test.expected)
		}
	}

	for _, test := range queueTests {
		fmt.Println(test.arg1)
		if myQueue.Dequeue(); len(myQueue.items) == -1 {
			fmt.Println("error")
			//t.Errorf("Output %q not equal to expected %q", len(stack.items), test.expected)
			t.Errorf("Output %d is not as expected %d", len(myQueue.items), test.expected)
		}
	}
}
