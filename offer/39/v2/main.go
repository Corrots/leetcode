package main

func main() {

}

// 摩尔投票法
// 初始化票数votes为0，众数为x
// 遍历数组中每个元素，当票数为0是，假设当前遍历到的元素x为众数，将遍历元素与x比较
// 若nums[i]=x，则票数+1，不相等则票数-1
// 循环结束返回x即为众数

// 时间复杂度：O(n)
// 空间复杂度：O(1)

//https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/submissions/
func majorityElement(nums []int) int {
	var x, votes int
	for _, num := range nums {
		if votes == 0 {
			x = num
			votes++
		} else {
			if x == num {
				votes++
			} else {
				votes--
			}
		}
	}
	return x
}
