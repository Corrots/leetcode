package main

func main() {

}

//
//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {

	}
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
