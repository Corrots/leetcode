package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{4, 9, 5}
	fmt.Println(getSum(nums))
}

//https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
func sumNumbers(root *TreeNode) int {
	var res [][]int
	var path []int
	getPaths(&res, path, root)
	var sum int
	for _, v := range res {
		sum += getSum(v)
	}
	return sum
}

func getSum(nums []int) int {
	var res int
	n := len(nums) - 1
	for i := 0; i <= n; i++ {
		if n == i {
			res += nums[i]
		} else {
			res += nums[i] * int(math.Pow10(n-i))
		}
	}
	return res
}

func getPaths(res *[][]int, path []int, root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		dst := make([]int, len(path)+1)
		copy(dst, append(path, root.Val))
		*res = append(*res, dst)
	} else {
		path = append(path, root.Val)
		getPaths(res, path, root.Left)
		getPaths(res, path, root.Right)
	}
}

//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
