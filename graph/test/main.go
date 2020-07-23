package main

import (
	"fmt"

	v1 "github.com/corrots/leetcode/graph/v1"
)

func main() {
	graph := v1.NewDG(4, false)
	fmt.Println(graph)
}
