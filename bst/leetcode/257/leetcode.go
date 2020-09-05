package main

import "strconv"

func main() {

}

//https://leetcode-cn.com/problems/binary-tree-paths/
func binaryTreePaths(root *TreeNode) []string {
	treePaths(root, "")
	return res
}

var res []string

func treePaths(root *TreeNode, path string) {
	if root == nil {
		return
	}
	subPath := path
	subPath += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		res = append(res, subPath)
	} else {
		subPath += "->"
		treePaths(root.Left, subPath)
		treePaths(root.Right, subPath)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
