package main

func main() {

}

// https://leetcode-cn.com/problems/sum-of-left-leaves/
func sumOfLeftLeaves(root *TreeNode) int {
	return sumLeftLeaves(root, false)
}

// 设置flag来标记是否为叶子结点
func sumLeftLeaves(root *TreeNode, isLeaf bool) int {
	if root == nil {
		return 0
	}
	if isLeaf && root.Left == nil && root.Right == nil {
		return root.Val
	}
	leftSum := sumLeftLeaves(root.Left, true)
	rightSum := sumLeftLeaves(root.Right, false)
	return leftSum + rightSum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
