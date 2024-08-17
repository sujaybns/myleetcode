package math

import (
	"fmt"
	"testing"
)

type  BFSTreeTest struct {
	arg1 int
	output int 
}

var BFSTreeTests = [] BFSTreeTest{
	{100, 0},
	{80,0},
	{110,0},
	{70,0},
}

func TestBFSTree(t *testing.T) {

	bfsTre := initBFSTree()
	for _ , test := range BFSTreeTests {
		if bfsTre.Insert(test.arg1); bfsTre.Search(test.arg1) != true {
			t.Errorf("Expected and actual values are different")
		} 
	}
	fmt.Println("BFS traversal")
	bfsTre.BFSTraversal()
	
}