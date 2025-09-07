package twopointers

import ("strings"
		"unicode")

func isPalindrome(s string) bool {
    strSpaceAndNumberFree := removeSpacesAndNumbers(s)
    uppercaseString := strings.ToUpper(strSpaceAndNumberFree)
    chars := []rune(uppercaseString)
    left :=0
    right := len(uppercaseString) -1
    for left < right {
        if chars[left] != chars[right] {
            return false
        }
        left++
        right--
    }
    return true
}
func isPalindromeDefault(s string) bool {
    strSpaceAndNumberFree := removeSpacesAndNumbers(s)
    uppercaseString := strings.ToUpper(strSpaceAndNumberFree)
    numberOfCharacters := len(uppercaseString)
    if numberOfCharacters == 1 {
        return  true
    }
    var isEven bool
    if numberOfCharacters%2 == 0{
        isEven = true
    }
    var myStack Stack
    chars := []rune(uppercaseString)
        for i :=0 ; i<numberOfCharacters/2 ; i++ {
            myStack.Push(chars[i])
        }
        startIndex := 0
        if isEven {
            startIndex = numberOfCharacters/2
        } else {
            startIndex = numberOfCharacters/2 + 1
        }
        for j :=startIndex; j < numberOfCharacters; j++{
            if !(myStack.Pop() == chars[j]){
                return false
            }
        }   
    return true
}
type StackInterface interface {
    Push(ch rune)
    Pop() rune
}

type Stack struct{
    stackArray [] rune
}

func (s *Stack) Push(ch rune){
    s.stackArray = append(s.stackArray,ch)
}

func (s *Stack) Pop() rune{
    lastIndex := len(s.stackArray) -1
    ch := s.stackArray[lastIndex]
    s.stackArray = s.stackArray[:lastIndex]
    return ch
}

func removeSpacesAndNumbers(s string) string {
    var result []rune
    for _, r := range s {
        // if !unicode.IsSpace(r) && !unicode.IsDigit(r) {
        if unicode.IsLetter(r) || unicode.IsDigit(r){
            result = append(result, r)
        }
    }
    return string(result)
}