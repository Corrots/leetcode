package main

func main() {

}

//https://leetcode-cn.com/problems/path-sum/
func hasPathSum(root *TreeNode, sum int) bool {
	result = make(map[int]bool)
	getSum(root, 0)
	if result[sum] {
		return true
	}
	return false
}

var result map[int]bool

func getSum(root *TreeNode, sub int) {
	//var result int
	if root != nil {
		toSum := sub
		toSum += root.Val
		if root.Left == nil && root.Right == nil {
			result[toSum] = true
		} else {
			getSum(root.Left, toSum)
			getSum(root.Right, toSum)
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
