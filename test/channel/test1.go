package main

import (
	"time"
)

// select 多个case同时满足时，会随机选择case执行
func main() {
	ch := make(chan int)
	go func() {
		for range time.Tick(time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")
		case <-ch:
			println("case2")
		}
	}
}
