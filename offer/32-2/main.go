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
	// i用于记录当前层数，从0开始
	for i := 0; len(q) > 0; i++ {
		// 定义一个slice来存储下一层的node
		var p []*TreeNode
		res = append(res, []int{})
		// 遍历当前q中的所有node
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
		// 将p中存储的下一层的node全部放入q中，继续下一轮循环
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
