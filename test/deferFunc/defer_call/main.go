package main

import (
	"fmt"
)

func main() {
	fmt.Println(defer_call())
}

func defer_call() int {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
	return 123
}
