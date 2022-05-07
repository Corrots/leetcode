package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	fmt.Println(nums1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/linked-list-cycle/
// 141. 环形链表
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	// 快慢指针
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
