package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s1))
	s2 := "race a car"
	fmt.Println(isPalindrome(s2))
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)
	n := len(s)
	i, j := 0, n-1
	for i < j {
		if !isValidChar(s[i]) {
			i++
			continue
		}
		if !isValidChar(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isValidChar(char byte) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
}
