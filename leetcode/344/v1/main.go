package main

import "fmt"

func main() {
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s1)
	fmt.Printf("%s\n", s1)
	s2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s2)
	fmt.Printf("%s\n", s2)
}

func reverseString(s []byte) {
	if s == nil {
		return
	}
	n := len(s)
	i, j := 0, n-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
