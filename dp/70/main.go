package main

import "fmt"

func main() {
	n := 10
	fmt.Println(climbStairs(n))
}

//爬楼梯
//https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	memo := make([]int, n+1)
	memo[0] = 1
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}
