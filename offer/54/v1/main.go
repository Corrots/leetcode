package main

func main() {

}

// 已知二叉树的中序遍历的数据是顺序的
// 先中序遍历保存到数组中，然后返回倒数第k个元素即为第k大的元素

// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
func kthLargest(root *TreeNode, k int) int {
	var order func(root *TreeNode)
	var res []int
	order = func(root *TreeNode) {
		if root == nil {
			return
		}
		order(root.Left)
		res = append(res, root.Val)
		order(root.Right)
		return
	}
	order(root)

	if k <= len(res) {
		return res[len(res)-k]
	}
	return 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
