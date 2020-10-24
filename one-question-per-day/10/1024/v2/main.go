package main

import (
	"fmt"
)

func main() {
	clips := [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	fmt.Println(videoStitching(clips, 10))
}

// 贪心算法求解
//https://leetcode-cn.com/problems/video-stitching/
func videoStitching(clips [][]int, T int) int {

}
