package findunion

import (
	"fmt"
	"testing"
)

func TestKruskals(t *testing.T) {
	edges := []Edge{
		{0, 1, 5},
		{0, 2, 10},
		{1, 2, 15}}
	value := kruskals(&edges)
	expected := 15
	if value != expected {
		t.Errorf("kruskals(&edges) = %d; want %d", value, expected)
	}
	fmt.Printf("Value:%v", value)
}
