package main

func main() {

}

// 已知二叉树的中序遍历的数据是升序的
// 则通过 right->root->left 的访问方式，就是降序的

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
func kthLargest(root *TreeNode, k int) int {
	steps := k
	var res int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		steps--
		if steps == 0 {
			res = root.Val
			return
		}
		dfs(root.Left)
	}
	dfs(root)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
