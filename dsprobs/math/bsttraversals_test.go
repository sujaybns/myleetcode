package math

import (
	"fmt"
	"slices"
	"testing"
)

type  BSTTraversalsTest struct {
	arg1 int
	output int 
}

var bstTraversalsTests = [] BSTTraversalsTest{
		{100,0},
		{80,0},
		{70,0},
		{60,0},
		{50,0},
		{40,0},
}

type BSTTraversalTreeTestOutput struct {
	output []int
}

var bstTraversalTreeTestInorderOutput = BSTTraversalTreeTestOutput{output: []int{40, 50, 60, 70, 80, 100}}
var bstTraversalTreeTestPreorderOutput = BSTTraversalTreeTestOutput{output: []int{100, 80, 70, 60, 50, 40}}
var bstTraversalTreeTestPostorderOutput = BSTTraversalTreeTestOutput{output: []int{40, 50, 60, 70, 80, 100}}

func TestTraversal(t *testing.T) {

	bstTre := initTraversalTree()
	for _ , test := range bstTraversalsTests {
		if bstTre.Insert(test.arg1); bstTre.Search(test.arg1) != true {
			t.Errorf("Expected and actual values are different")
		} 
	}
	inOrderTraversal := []int{}
	if bstTre.InOrderTraversal(&inOrderTraversal); slices.Equal(bstTraversalTreeTestInorderOutput.output, inOrderTraversal) != true {
		t.Error("Expected and actual values are different")
	}
	fmt.Println(bstTraversalTreeTestInorderOutput)
	preOrderTraversal := []int{}
	if bstTre.PreOrderTraversal(&preOrderTraversal); slices.Equal(bstTraversalTreeTestPreorderOutput.output, preOrderTraversal) != true {
		t.Error("Expected and actual values are different")
	}
	fmt.Println(bstTraversalTreeTestPreorderOutput)
	postOrderTraversal := []int{}
	if bstTre.PostOrderTraversal(&postOrderTraversal); slices.Equal(bstTraversalTreeTestPostorderOutput.output, postOrderTraversal) != true {
		t.Error("Expected and actual values are different")
	}
	fmt.Println(bstTraversalTreeTestPostorderOutput)
}