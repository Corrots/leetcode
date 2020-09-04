package main

func main() {

}

//https://leetcode-cn.com/problems/path-sum/
func hasPathSum(root *TreeNode, sum int) bool {
	result = []int{}
	getSum(root, 0)
	for _, v := range result {
		if v == sum {
			return true
		}
	}
	return false
}

var result []int

func getSum(root *TreeNode, sub int) {
	//var result int
	if root != nil {
		toSum := sub
		toSum += root.Val
		if root.Left == nil && root.Right == nil {
			result = append(result, toSum)
		} else {
			getSum(root.Left, toSum)
			getSum(root.Right, toSum)
		}
	}
}
