package twopointers

import "testing"

func TestThreeSum_Optimized(t *testing.T) {
	triplets := threeSumOptimized([]int{-1,0,1,0})
	if !slicesEqualSet(triplets[0],[]int{-1,0,1}) {
		t.Error("Output is not the same")
	}
}

func TestThreeSum_Default(t *testing.T) {
	triplets := threeSumOptimized([]int{-1,0,1,0})
	if !slicesEqualSet(triplets[0],[]int{-1,0,1}) {
		t.Error("Output is not the same")
	}
}