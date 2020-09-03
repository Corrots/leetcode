package main

import (
	"fmt"
)

func main() {
	s1 := "hello"
	fmt.Printf("%s\n", reverseVowels(s1))
	s2 := "leetcode"
	fmt.Printf("%s\n", reverseVowels(s2))
	s3 := "aA"
	fmt.Printf("%s\n", reverseVowels(s3))
}

func reverseVowels(s string) string {
	n := len(s)
	bytes := []byte(s)
	i, j := 0, n-1
	for i < j {
		if isVowel(bytes[i]) && isVowel(bytes[j]) {
			bytes[i], bytes[j] = bytes[j], bytes[i]
			i++
			j--
			continue
		}
		if !isVowel(bytes[i]) {
			i++
			continue
		}
		if !isVowel(bytes[j]) {
			j--
		}
	}
	return string(bytes)
}

func isVowel(char uint8) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U'
}
