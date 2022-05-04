package main

import "fmt"

type MySlice[T int | float64] []T

func (s MySlice[T]) Sum() T {
	var total T
	for _, val := range s {
		total += val
	}
	return total
}

func main() {
	var a MySlice[int] = []int{1, 2, 3}
	fmt.Println(a.Sum())

	var b MySlice[float64] = []float64{0.1, 0.2, 0.3}
	fmt.Println(b.Sum())
}
