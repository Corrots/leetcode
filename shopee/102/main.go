package main

/**
引入队列对二叉树进行层序遍历

*/

//https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return nil
	}
	// 将根节点加入队列中
	q := []*TreeNode{root}
	// 分别遍历每一层，i代表当前层数
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		// 队列p用于存放每一层的节点
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		// 将p赋值给q，即将当前层遍历到的节点全部放入外层队列中，进一步遍历其孩子节点
		q = p
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
