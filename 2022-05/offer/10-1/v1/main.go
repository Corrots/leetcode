package main

import "fmt"

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(5))
}

// https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof/
// 斐波拉契数列
//
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
