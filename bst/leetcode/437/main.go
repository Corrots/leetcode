package main

func main() {

}

//https://leetcode-cn.com/problems/path-sum-iii/
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := findPath(root, sum)
	return res + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func findPath(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var res int
	if sum == root.Val {
		res += 1
	}
	return res + findPath(root.Left, sum-root.Val) + findPath(root.Right, sum-root.Val)
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
