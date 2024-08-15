package math

import (
	"fmt"
	"testing"
)

type  BSTTreeTest struct {
	arg1 int
	output int 
}

var bstTreeTests = [] BSTTreeTest{
	{100, 0},
	{80,0},
	{110,0},
	{70,0},
}

func TestBSTTree(t *testing.T) {

	bstTre := initBSTTree()
	for _ , test := range bstTreeTests {
		if bstTre.Insert(test.arg1); bstTre.Search(test.arg1) != true {
			t.Errorf("Expected and actual values are different")
		} 
	}
	for _ , test := range bstTreeTests {
		if bstTre.Delete(test.arg1); bstTre.Search(test.arg1) == true {
			t.Errorf("Expected and actual values are different")
		} 
	}
}

func TestBSTTreeAllOperations(t *testing.T) {

	bstTre := initBSTTree()
	bstTre.Insert(100)
	bstTre.Insert(80)
	bstTre.Insert(110)
	bstTre.Insert(70)
	// fmt.Printf("bstTre.Search(70):%v\n", bstTre.Search(70))
	bstTre.inOrderTraversal(bstTre.root)
	bstTre.Delete(80)
	bstTre.inOrderTraversal(bstTre.root)
	// bstTre.LeftSearch()
	fmt.Println("Complete")
}
