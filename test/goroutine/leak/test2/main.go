package main

import (
	"fmt"
	"runtime"
)

func main() {
	defer func() {
		//time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	//var ch chan struct{}
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
		defer close(ch) // 完成发送后及时关闭channel
	}()
	<-ch
}
