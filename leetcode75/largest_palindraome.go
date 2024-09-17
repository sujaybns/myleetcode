package leetcode75

func expandFromCenter(i, j int, mystring string) string {
	var visited bool
	for i >= 0 && j < len(mystring) && mystring[i] == mystring[j] {
		visited = true
		i = i - 1
		j = j + 1
	}
	if visited {
		//Learning: 
		//str := "Hello, Go!" => str[0:5] produces the substring "Hello" 
		//because it includes characters from index 0 to index 4 
		//(H(0), e(1), l(2), l(3), o(4)).
		return mystring[i+1 : j]
	}
	return ""
}

func getLongestPalindrome(mystring string) string {
	var val string
	for i := 0; i < len(mystring); i++ {
		newVal := expandFromCenter(i, i+1, mystring)
		if len(val) < len(newVal) {
			val = newVal
		}
		newOddVal := expandFromCenter(i, i, mystring)
		if len(val) < len(newOddVal) {
			val = newOddVal
		}
	}
	return val
}
