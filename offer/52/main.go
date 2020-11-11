package main

func main() {

}

// 定义p，q分别指向两个链表的head，遍历链表
// 当p走到headA的末尾时，将其重新指向HeadB
// 当q走到headB的末尾时，将其重新指向HeadA
// 最后当 p == q时，即找到两个链表的公共节点

//https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p, q := headA, headB
	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}
		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}
	return p
}

//
type ListNode struct {
	Val  int
	Next *ListNode
}
