package leetcode75

func expandFromCenter(i, j int, mystring string) string {
	var visited bool
	for i >= 0 && j < len(mystring) && mystring[i] == mystring[j] {
		visited = true
		i = i - 1
		j = j + 1
	}
	if visited {
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
