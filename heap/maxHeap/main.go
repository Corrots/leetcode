package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/corrots/leetcode/heap/maxHeap/heap"
)

func main() {
	maxHeap := heap.New(100)

	rand.Seed(time.Now().UnixNano())
	//arr := []int{62, 41, 30, 28, 16, 22, 13, 19, 17, 15}
	//for _, v := range arr {
	//	//maxHeap.Insert(rand.Int() % 100)
	//	maxHeap.Insert(v)
	//}
	for i := 0; i < 20; i++ {
		maxHeap.Insert(rand.Int() % 100)
	}
	maxHeap.Insert(52)
	fmt.Println(maxHeap)
	for !maxHeap.IsEmpty() {
		fmt.Printf("%d ", maxHeap.ExtractMax())
	}
}
