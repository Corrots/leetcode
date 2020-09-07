package main

func main() {

}

//https://leetcode-cn.com/problems/path-sum-iii/
// 在以root为根节点的二叉树中寻找和为sum的路径个数
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := findPath(root, sum)
	return res + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

// 在以node为根节点的二叉树中，寻找从node开始的路径，且路径的和为sum的个数
func findPath(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	var res int
	if sum == node.Val {
		res += 1
	}
	return res + findPath(node.Left, sum-node.Val) + findPath(node.Right, sum-node.Val)
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
