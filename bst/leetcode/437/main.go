package main

func main() {

}

// 包含 当前节点在路径中 或 当前节点不在路径中两种场景
// findPath(root) + pathSum(root.left, sum) + pathSum()

//https://leetcode-cn.com/problems/path-sum-iii/
// 在以root为根节点的二叉树中寻找和为sum的路径个数
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var res int
	// 当前节点包含在路径中时
	res = findPath(root, sum)
	res += pathSum(root.Left, sum)
	res += pathSum(root.Right, sum)
	return res
}

// 在以node为根节点的二叉树中，寻找node节点包含在路径中，切和为num的个数
func findPath(node *TreeNode, num int) int {
	if node == nil {
		return 0
	}
	var res int
	if node.Val == num {
		res++
	}
	res += findPath(node.Left, num-node.Val)
	res += findPath(node.Right, num-node.Val)
	return res
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
