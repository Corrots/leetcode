package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	memo := make([]int, x+1)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i < x+1; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[x]
}
