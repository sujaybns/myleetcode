package twopointers

import (
	"testing"
	"sort"
)

func TestTwoSum(t *testing.T) {
	pair := TwoSum([]int{1, 2, 3}, 5)
	if !slicesEqualSet(pair,[]int{1,2}) {
		t.Error("Output is not the same")
	}
}

func TestTwoDefault(t *testing.T) {
	pair := TwoSumDefault([]int{1, 2, 3}, 5)
	if !slicesEqualSet(pair,[]int{1,2}) {
		t.Error("Output is not the same")
	}
}


func slicesEqualSet(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    aCopy := append([]int(nil), a...)
    bCopy := append([]int(nil), b...)
    sort.Ints(aCopy)
    sort.Ints(bCopy)
    for i := range aCopy {
        if aCopy[i] != bCopy[i] {
            return false
        }
    }
    return true
}