package main

import "fmt"

// https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/
func main() {
	str := "abbaca"
	fmt.Println(removeDuplicates(str) == "ca")
}

func removeDuplicates(s string) string {
	var stack []byte
	for i := range s {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
