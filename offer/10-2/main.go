package main

import "fmt"

func main() {
	fmt.Println(numWays(2))
	fmt.Println(numWays(3))
}

//https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	memo := make([]int, n+1)
	memo[0] = 1
	memo[1] = 1
	for i := 2; i < n+1; i++ {
		memo[i] = (memo[i-1] + memo[i-2]) % (1e9 + 7)
	}
	return memo[n]
}
