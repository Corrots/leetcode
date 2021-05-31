package main

//https://leetcode-cn.com/problems/binary-tree-right-side-view/
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		res = append(res, q[n-1].Val)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
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
