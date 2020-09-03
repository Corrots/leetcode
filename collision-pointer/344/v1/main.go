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
	if len(s) == 0 {
		return
	}
	p, q := 0, len(s)-1
	for p < q {
		s[p], s[q] = s[q], s[p]
		p++
		q--
	}
}
