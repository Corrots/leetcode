package main

func main() {

}

//https://leetcode-cn.com/problems/path-sum-ii/
func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	var path []int
	sumPath(&res, path, root, sum)
	return res
}

func sumPath(res *[][]int, path []int, root *TreeNode, sum int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		dst := make([]int, len(path)+1)
		copy(dst, append(path, root.Val))
		*res = append(*res, dst)
	} else {
		path = append(path, root.Val)
		sumPath(res, path, root.Left, sum-root.Val)
		sumPath(res, path, root.Right, sum-root.Val)
	}
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
