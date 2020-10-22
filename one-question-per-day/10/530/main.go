package main

import (
	"math"
)

func main() {

}

//          4
//        /   \
//      2      6
//     / \
//    1   3

// BST的中序遍历是有序的，比较相邻元素的差值，获取最小值
//https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
func getMinimumDifference(root *TreeNode) int {
	res := math.MaxInt64
	prev := -1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != -1 && node.Val-prev < res {
			res = node.Val - prev
		}
		prev = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return res
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
