package math

import "testing"

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type pushTest struct {
	arg1, expected int
}

var pushTests = []pushTest{
	pushTest{2, 1},
	pushTest{4, 2},
	pushTest{6, 3},
	pushTest{3, 4},
}

func TestPush(t *testing.T) {
	stack := Stack{}
	for _, test := range pushTests {

		if stack.Push(test.arg1); len(stack.items) != test.expected {
			t.Errorf("Output %q not equal to expected %q", len(stack.items), test.expected)
		}
	}

	for _, test := range pushTests {

		if stack.Pop(); len(stack.items) == -1 {
			t.Errorf("Output %q not equal to expected %q", len(stack.items), test.expected-1)
		}
	}
}
:q:q
