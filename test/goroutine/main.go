package main

import "fmt"

// 通过channel，让goroutine顺序执行
func main() {
	ch := make(chan int)
	for i := 0; i < 6; i++ {
		go func(ch <-chan int) {
			fmt.Println(<-ch)
		}(ch)
		ch <- i
	}
}
