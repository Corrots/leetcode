package main

import "fmt"

// Slice 范型
type Slice[T int | float32 | float64] []T

func main() {
	var a Slice[int] = []int{1, 2, 3}
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var b Slice[float64] = []float64{0.1, 0.2, 0.3}
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	var c Slice[float32] = []float32{0.01, 0.02, 0.03}
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
