package leetcode75

import "testing"

func TestLP(t *testing.T) {
	if pal := getLongestPalindrome("abba"); pal != "abba" {
		t.Error("Test failed")
	}
	if pal := getLongestPalindrome("aba"); pal != "aba" {
		t.Error("Test failed")
	}
}
