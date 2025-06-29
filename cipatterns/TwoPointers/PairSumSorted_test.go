package twopointers

import (
	"fmt"
	"testing"
)

func TestPairSum(t *testing.T) {
	pair := PairSumDefault(5, []int{1, 2, 3})
	fmt.Println(pair.SIndex, pair.LIndex)
	pair = PairSumOptimized(5, []int{1, 2, 3})
	fmt.Println(pair.SIndex, pair.LIndex)
}
