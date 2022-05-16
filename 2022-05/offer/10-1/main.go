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
	if n <= 1 {
		return n
	}
	memo := make([]int, n+1)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = (memo[i-1] + memo[i-2]) % (1e9 + 7)
	}
	return memo[n]
}
