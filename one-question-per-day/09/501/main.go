package main

import "math"

func main() {

}

//https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/
func findMode(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var inOrder func(*TreeNode)
	var list []int
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		list = append(list, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	var count, maxCount int
	base := math.MinInt64
	for i := 0; i < len(list); i++ {
		if base == list[i] {
			count++
		} else {
			base = list[i]
			count = 1
		}
		if count == maxCount {
			res = append(res, base)
		} else if count > maxCount {
			maxCount = count
			res = []int{}
			res = append([]int{}, base)
		}
	}
	return res
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
