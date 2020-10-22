package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums[1:])
}

// 尝试使用O(1)的空间复杂度解决
//https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	return nil
}

//
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
