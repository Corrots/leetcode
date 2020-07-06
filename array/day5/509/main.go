package main

import "fmt"

/**
#509
斐波那契数
斐波那契数，通常用 F(n) 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。
*/
func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(7))
}

func fib(N int) int {
	if N == 0 {
		return 0
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
