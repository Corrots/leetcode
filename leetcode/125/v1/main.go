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
	var chars string
	for i := 0; i < len(s); i++ {
		if isValidChar(s[i]) {
			chars += string(s[i])
		}
	}
	chars = strings.ToLower(chars)
	n := len(chars)
	l, r := 0, n-1
	for l < r {
		if chars[l] == chars[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func isValidChar(char byte) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
}
