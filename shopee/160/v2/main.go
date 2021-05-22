package main

/**
双指针
*/

//https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	head1, head2 := headA, headB
	for head1 != head2 {
		if head1 != nil {
			head1 = head1.Next
		} else {
			head1 = headB
		}

		if head2 != nil {
			head2 = head2.Next
		} else {
			head2 = headA
		}
	}
	return head1
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
