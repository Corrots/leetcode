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
func fib(N int) int {
	if N <= 1 {
		return N
	}
	var res []int
	fn := getFib()
	for i := 0; i < N; i++ {
		res = append(res, fn())
	}

	return res[N-1]
}

func getFib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
