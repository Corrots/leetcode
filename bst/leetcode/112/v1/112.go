package main

func main() {

}

//https://leetcode-cn.com/problems/path-sum/
func hasPathSum(root *TreeNode, sum int) bool {
	result := make(map[int]bool)
	getSum(&result, root, 0)
	if result[sum] {
		return true
	}
	return false
}

func getSum(result *map[int]bool, root *TreeNode, sub int) {
	if root != nil {
		toSum := sub
		toSum += root.Val
		if root.Left == nil && root.Right == nil {
			(*result)[toSum] = true
		} else {
			getSum(result, root.Left, toSum)
			getSum(result, root.Right, toSum)
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
