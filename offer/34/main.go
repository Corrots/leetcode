package main

func main() {

}

// 深度优先遍历

//https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	var dfs func(*TreeNode, int, []int)
	dfs = func(node *TreeNode, sum int, path []int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && node.Val == sum {
			res = append(res, append(path, node.Val))
			return
		}
		tmp := make([]int, len(path))
		copy(tmp, path)
		tmp = append(tmp, node.Val)
		dfs(node.Left, sum-node.Val, tmp)
		dfs(node.Right, sum-node.Val, tmp)
	}
	dfs(root, sum, []int{})
	return res
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
