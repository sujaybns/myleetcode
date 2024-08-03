package math

import (
	"fmt"
	"testing"
)

type LinkedlistTest struct {
	arg1, expected int
}

var LinklistTests = []LinkedlistTest{
	LinkedlistTest{1, 1},
	LinkedlistTest{10, 2},
	LinkedlistTest{5, 3},
	LinkedlistTest{50, 4},
}

func TestLinkedListInsertion(t *testing.T) {
	mylinkedlist := Linkedist{}
	for _, test := range LinklistTests {
		if mylinkedlist.Insert(&Node{val: test.arg1}); mylinkedlist.length != test.expected {
			fmt.Println("error")
			t.Errorf("Output %d is not as expected %d", mylinkedlist.length, test.expected)
		}
	}
}

func TestLinkedListDeletion(t *testing.T) {
	mylinkedlist := Linkedist{}
	for _, test := range LinklistTests {
		if mylinkedlist.Insert(&Node{val: test.arg1}); mylinkedlist.length != test.expected {
			fmt.Println("error")
			t.Errorf("Output %d is not as expected %d", mylinkedlist.length, test.expected)
		}
	}

	for _, test := range LinklistTests {
		initLength := mylinkedlist.length
		if mylinkedlist.Delete(&Node{val: test.arg1}); mylinkedlist.length != initLength-1 {
			fmt.Println("error")
			t.Errorf("Output %d is not as expected %d", mylinkedlist.length, initLength-1)
		}
	}
}
