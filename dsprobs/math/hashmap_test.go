package math

import (
	"testing"
)

type HashmapTest struct {
	arg1   string
	output int
}

var hashMapTests = []HashmapTest{
	HashmapTest{arg1: "sujay", output: 1},
	HashmapTest{arg1: "john", output: 2},
	HashmapTest{arg1: "mary", output: 3},
}

func TestHashMapInsertTest(t *testing.T) {
	hashTable := Init()
	for _, test := range hashMapTests {
		if hashTable.Insert(test.arg1); hashTable.Search(test.arg1) != true {
			t.Error("Expected value and the actual value do not match")
		}
	}
	for _, test := range hashMapTests {
		if hashTable.Delete(test.arg1); hashTable.Search(test.arg1) != false {
			t.Error("Expected value and the actual value do not match")
		}
	}
}
