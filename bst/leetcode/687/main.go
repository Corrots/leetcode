package main

func main() {

}

//https://leetcode-cn.com/problems/longest-univalue-path/
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := findPath(root, root.Val) - 1
	return max(res, longestUnivaluePath(root.Left)+longestUnivaluePath(root.Right))
}

func findPath(node *TreeNode, val int) int {
	if node == nil {
		return 0
	}
	var res int
	if node.Val == val {
		res++
	}
	return res + findPath(node.Left, node.Val) + findPath(node.Right, node.Val)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
