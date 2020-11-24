package main

func main() {

}

// 完全二叉树，叶子节点只会出现在最后两层

//https://leetcode-cn.com/problems/count-complete-tree-nodes/
func countNodes(root *TreeNode) int {
	var res int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		res++
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
