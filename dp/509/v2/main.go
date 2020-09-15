package main

import (
	"fmt"
	"time"
)

func main() {
	n := 50
	start := time.Now()
	res := fib(n)
	fmt.Println(res)
	fmt.Printf("spend %d ms\n", time.Since(start).Milliseconds())
}

//https://leetcode-cn.com/problems/fibonacci-number/
// 动态规划 - 自上向下的解决问题
func fib(N int) int {
	if N <= 1 {
		return N
	}
	memo := make([]int, N+1)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= N; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[N]
}
