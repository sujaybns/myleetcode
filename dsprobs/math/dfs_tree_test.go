package math

import (
	"fmt"
	"slices"
	"testing"
)

type DFSBSTTreeTest struct {
	input int
}

var dfsBSTTreeTests = []DFSBSTTreeTest{
	{100},
	{80},
	{70},
	{60},
	{50},
	{40},
}

type DFSBSTTreeTestOutput struct {
	output []int
}

var dfsBSTTreeTestOutput = DFSBSTTreeTestOutput{output: []int{100, 80, 70, 60, 50, 40}}

func TestDFS(t *testing.T) {
	dfsBSTTree := initDFSBSTTree()
	for _, test := range dfsBSTTreeTests {
		if dfsBSTTree.DFSInsert(test.input); dfsBSTTree.DFSSearch(test.input) != true {
			t.Error("Expected and actual values are different")
		}
	}
	traversal := []int{}
	if dfsBSTTree.DFSTraversal(&traversal); slices.Equal(dfsBSTTreeTestOutput.output, traversal) != true {
		t.Error("Expected and actual values are different")
	}
	fmt.Println(traversal)
}
