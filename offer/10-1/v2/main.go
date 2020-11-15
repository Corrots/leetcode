package main

import "fmt"

func main() {
	fmt.Println(fib(45))
	fmt.Printf("counter: %d\n", counter)
}

var counter int

// 动态规划求解

//https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
func fib(n int) int {
	if n <= 1 {
		return n
	}
	memo := make([]int, n+1)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i < n+1; i++ {
		memo[i] = (memo[i-1] + memo[i-2]) % (1e9 + 7)
	}
	return memo[n]
}
