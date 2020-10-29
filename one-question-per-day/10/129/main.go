package main

func main() {

}

// 通过深度优先搜索求解
// 1.如果当前节点是叶子节点，直接将当前节点的值加到数字之和
// 2.如果当前节点是非叶子结点，遍历其叶子节点

//https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, preSum int) int {
	if root == nil {
		return 0
	}
	sum := preSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
