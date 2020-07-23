package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/corrots/leetcode/unionfind/uf2"
	"github.com/corrots/leetcode/unionfind/uf3"
	"github.com/corrots/leetcode/unionfind/uf4"
)

func main() {
	n := 1000000
	TestUF3(n)
	TestUF4(n)
}

func TestUF2(n int) {
	rand.Seed(time.Now().UnixNano())
	uf := uf2.New(n)
	start := time.Now()

	for i := 0; i < n; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		uf.Union(a, b)
	}
	for i := 0; i < n; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		uf.IsConnected(a, b)
	}
	fmt.Printf("UF2, %d ops, %d ms\n", 2*n, time.Since(start).Milliseconds())
}

func TestUF3(n int) {
	rand.Seed(time.Now().UnixNano())
	uf := uf3.New(n)
	start := time.Now()

	for i := 0; i < n; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		uf.Union(a, b)
	}
	for i := 0; i < n; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		uf.IsConnected(a, b)
	}
	fmt.Printf("UF3, %d ops, %d ms\n", 2*n, time.Since(start).Milliseconds())
}

func TestUF4(n int) {
	rand.Seed(time.Now().UnixNano())
	uf := uf4.New(n)
	start := time.Now()

	for i := 0; i < n; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		uf.Union(a, b)
	}
	for i := 0; i < n; i++ {
		a := rand.Int() % n
		b := rand.Int() % n
		uf.IsConnected(a, b)
	}
	fmt.Printf("UF4, %d ops, %d ms\n", 2*n, time.Since(start).Milliseconds())
}
