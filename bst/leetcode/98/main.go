package main

import (
	"math"
)

func main() {

}

//https://leetcode-cn.com/problems/validate-binary-search-tree/
// 检查以root为根节点的二叉树是否为二叉搜索树
func isValidBST(root *TreeNode) bool {
	return checkBST(root, math.MinInt64, math.MaxInt64)
}

// 检查以node为根节点的二叉树是否为二叉搜索树
func checkBST(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lower || node.Val >= upper {
		return false
	}
	return checkBST(node.Left, lower, node.Val) && checkBST(node.Right, node.Val, upper)
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
