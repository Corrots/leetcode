package main

import "strconv"

func main() {

}

//https://leetcode-cn.com/problems/binary-tree-paths/
func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		res = append(res, strconv.Itoa(root.Val))
		return res
	}
	for _, v := range binaryTreePaths(root.Left) {
		res = append(res, strconv.Itoa(root.Val)+"->"+v)
	}
	for _, v := range binaryTreePaths(root.Right) {
		res = append(res, strconv.Itoa(root.Val)+"->"+v)
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
