package findunion

import (
	"fmt"
	"testing"
)

func TestMake(t *testing.T) {
	parent := make([]int, 10)
	fmt.Printf("Before - parent : %v", parent)
	initParent(&parent)
	fmt.Printf("After - parent : %v", parent)
	fmt.Printf("find(2):%v", find(2, &parent))
}

func TestFind(t *testing.T) {
	parent := make([]int, 10)
	fmt.Printf("Before - parent : %v", parent)
	initParent(&parent)
	fmt.Printf("After - parent : %v", parent)
	fmt.Printf("find(2):%v", find(2, &parent))
}

func TestUnion(t *testing.T) {
	parent := make([]int, 10)
	fmt.Printf("Before - parent : %v", parent)
	initParent(&parent)
	fmt.Printf("After - parent : %v", parent)
	union(1,2,&parent)
	fmt.Printf("find(2):%v", find(2, &parent))
}