package main

import "fmt"

func main() {
	fmt.Println(fib(10))
	fmt.Printf("counter: %d\n", counter)
}

var counter int

//https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
func fib(n int) int {
	counter++
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// 记忆化搜索，记录斐波那契数从[0...n]
	memo := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = -1
	}
	if memo[n] == -1 {
		memo[n] = fib(n-1) + fib(n-2)
	}
	return memo[n]
}
