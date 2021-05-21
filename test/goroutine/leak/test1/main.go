package main

import (
	"fmt"
	"runtime"
	"time"
)

// channel只发送不接收，或者没有接收完全，都会导致发送方阻塞，进而引起goroutine泄漏
func gen(done chan bool, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	done := make(chan bool)
	defer close(done)
	// Set up the pipeline.
	out := gen(done, 2, 3)

	for n := range out {
		fmt.Println(n)              // 2
		time.Sleep(3 * time.Second) // done thing, 可能异常中断接收
		if true {                   // if err != nil
			break
		}
	}
}
