package main

import "math"

func main() {

}

//https://leetcode-cn.com/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode) bool {
	return checkBST(root, math.MinInt64, math.MaxInt64)
}

func checkBST(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return checkBST(root.Left, lower, root.Val) && checkBST(root.Right, root.Val, upper)
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
