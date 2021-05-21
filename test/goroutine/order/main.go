package main

import (
	"fmt"
)

// 两个goroutine, 一个goroutine输出1,3,5,7,9, 另一个输出2,4,6,8,10, 怎样有顺序的输出1,2,3,4,5,6,7,8
func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	done := make(chan bool)

	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println(2*i - 1)
			ch2 <- true
			<-ch1
		}
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			<-ch2
			fmt.Println(2 * i)
			ch1 <- true
		}
		defer func() { close(done) }()
	}()
	<-done
}
