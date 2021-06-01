package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("main process")
	go handle()

	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine())
}

func handle() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	for i := 0; i < 5; i++ {
		if i == 3 {
			panic("err")
		}
	}
}
