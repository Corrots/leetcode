package main

func main() {

}

//https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	if root == nil {
		return res
	}
	var nodes [][]int
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		var p []*TreeNode
		nodes = append(nodes, []int{})
		for j := 0; j < len(q); j++ {
			node := q[j]
			nodes[i] = append(nodes[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	for i := 0; i < len(nodes); i++ {
		res = append(res, average(nodes[i]))
	}
	return res
}

func average(args []int) float64 {
	var sum int
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return float64(sum) / float64(len(args))
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
