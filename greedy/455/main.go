package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	fmt.Println(findContentChildren(g, s))
}

// https://leetcode-cn.com/problems/assign-cookies/
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gi, si := len(g)-1, len(s)-1
	var res int
	for gi >= 0 && si >= 0 {
		if s[si] >= g[gi] {
			res++
			si--
		}
		gi--
	}
	return res
}
