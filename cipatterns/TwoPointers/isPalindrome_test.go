package twopointers

import "testing"

func TestIsPalindromeDefault(t *testing.T){
	result := isPalindromeDefault("0A")
	if result == true {
		t.Errorf("The isPlaindrome value for the number %s is not %t", "0A", result)
	}
	result = isPalindromeDefault("BAB")
	if result == false {
		t.Errorf("The isPlaindrome value for the number %s is not %t", "0A", result)
	}
}

func TestIsPalindrome(t *testing.T){
	result := isPalindrome("0A")
	if result == true {
		t.Errorf("The isPlaindrome value for the number %s is not %t", "0A", result)
	}
	result = isPalindrome("BAB")
	if result == false {
		t.Errorf("The isPlaindrome value for the number %s is not %t", "0A", result)
	}
}