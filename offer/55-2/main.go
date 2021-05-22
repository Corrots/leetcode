package main

//https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/
func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

// 求node的高度
func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := height(node.Left)
	rightHeight := height(node.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
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

func main() {

}
