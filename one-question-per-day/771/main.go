package main

import "fmt"

func main() {
	J := "aA"
	S := "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))
	fmt.Println(numJewelsInStones("z", "ZZ"))
}

//https://leetcode-cn.com/problems/jewels-and-stones/
func numJewelsInStones(J string, S string) int {
	var res int
	if len(S) == 0 {
		return res
	}
	jewels := make(map[int32]bool)
	for _, v := range J {
		jewels[v] = true
	}
	for _, v := range S {
		if jewels[v] {
			res++
		}
	}
	return res
}
