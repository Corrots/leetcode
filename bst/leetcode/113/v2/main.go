package main

func main() {

}

//https://leetcode-cn.com/problems/path-sum-ii/
func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var dfs func(*TreeNode, int, []int)
	dfs = func(node *TreeNode, sum int, path []int) {
		if node == nil {
			return
		}
		// 递归终止条件：当前节点为叶子节点，且sum == 当前节点的Val
		if node.Left == nil && node.Right == nil && sum == node.Val {
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
