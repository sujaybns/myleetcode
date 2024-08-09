package math

import (
          "testing"
		)

type TrieTest struct{
	arg1 string
	output int
} 

var trieTests = [] TrieTest {
	{arg1:"john", output:1},
	{arg1:"joseph", output:2},
}

func TestTrieInsert(t *testing.T){
	myTrie := initTrie()
	for _,test := range trieTests {
		if myTrie.TrieInsert(test.arg1);myTrie.TrieSearch(test.arg1) == false{
			t.Errorf("expected and actual do not match ")
		}
	}
}