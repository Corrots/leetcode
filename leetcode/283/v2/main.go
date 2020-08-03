package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

// 时间复杂度: O(n)
// 空间复杂度: O(1)
func moveZeroes(nums []int) {
	// nums[0:k)中的元素都是非0元素
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != k { //
				nums[k], nums[i] = nums[i], nums[k]
			}
			k++
		}
	}
}
