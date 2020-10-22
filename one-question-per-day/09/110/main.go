package main

func main() {

}

//一棵高度平衡二叉树定义为：一个二叉树每个节点的左右两个子树的高度差的绝对值不超过1。
//https://leetcode-cn.com/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	return height(root) == -1
}

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
		return -1 * i
	}
	return i
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
