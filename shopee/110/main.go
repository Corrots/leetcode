package main

// https://leetcode-cn.com/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := height(root.Left)
	right := height(root.Right)
	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}
	return max(left, right) + 1
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
