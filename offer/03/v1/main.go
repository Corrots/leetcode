package main

func main() {

}

// 思路：利用哈希表记录元素，即查找表
//https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
func findRepeatNumber(nums []int) int {
	existed := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if existed[nums[i]] {
			return nums[i]
		} else {
			existed[nums[i]] = true
		}
	}
	return 0
}
