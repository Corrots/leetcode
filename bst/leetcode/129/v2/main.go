package main

func main() {

}

//https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
func sumNumbers(root *TreeNode) int {
	return sum(root, 0)
}

func sum(root *TreeNode, i int) int {
	if root == nil {
		return 0
	}
	tmp := i*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return tmp
	}
	return sum(root.Left, tmp) + sum(root.Right, tmp)
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
