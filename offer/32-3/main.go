package main

func main() {

}

//https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		var p []*TreeNode
		n := len(q)
		row := make([]int, n)
		for j := 0; j < n; j++ {
			node := q[j]
			if i%2 == 0 {
				// 偶数行
				row[j] = node.Val
			} else {
				// 奇数行
				row[n-1-j] = node.Val
			}
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		res = append(res, row)
		q = p
	}
	return res
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
