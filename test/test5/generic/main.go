package main

import "fmt"

type MyMap[Key int | string, Value float32 | float64] map[Key]Value

func main() {
	newMap := make(MyMap[int, float64])
	newMap[1] = 0.1
	newMap[2] = 0.2

	fmt.Printf("val: %v, Type: %T\n", newMap, newMap)
}
