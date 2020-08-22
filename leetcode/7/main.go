package main

import "fmt"

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}

// https://leetcode-cn.com/problems/reverse-integer/
func reverse(x int) int {
	var res int
	for x != 0 {
		if tmp := int32(res); (tmp*10)/10 != tmp {
			return 0
		}
		res = res*10 + x%10
		x = x / 10
	}
	return res
}
