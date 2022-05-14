package main

func main() {

}

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
// 剑指 Offer 32 - III. 从上到下打印二叉树 III
// 二叉树的Z字形遍历，输出为二维数组
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		var p []*TreeNode
		var rows []int

		for j := 0; j < len(q); j++ {
			node := q[j]
			rows = append(rows, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		// 将奇数层元素翻转
		if i%2 == 1 {
			n := len(rows)
			for k := 0; k < n/2; k++ {
				rows[k], rows[n-1-k] = rows[n-1-k], rows[k]
			}
		}

		res = append(res, rows)
		q = p
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
