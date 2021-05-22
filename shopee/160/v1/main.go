package main

/**
HashMap求解
O(m+n)
*/

//https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	m := make(map[*ListNode]bool)
	cur := headA
	for cur != nil {
		m[cur] = true
		cur = cur.Next
	}
	cur = headB
	for cur != nil {
		if m[cur] {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
