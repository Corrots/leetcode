package main

func main() {

}

/**
快慢指针
*/

// https://leetcode-cn.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
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

type ListNode struct {
	Val  int
	Next *ListNode
}
