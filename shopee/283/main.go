package main

func main() {

}

//https://leetcode-cn.com/problems/move-zeroes/
func moveZeroes(nums []int) {
	if len(nums) <= 0 {
		return
	}
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}
	//
	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}
