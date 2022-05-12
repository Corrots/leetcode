package main

import "fmt"

func main() {
	q := Constructor()
	q.AppendTail(1)
	q.AppendTail(2)
	q.AppendTail(3)
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
	q.AppendTail(4)
	fmt.Println(q.DeleteHead())
}

//https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
// 剑指 Offer 09. 用两个栈实现队列

type CQueue struct {
	inStack  []int // 用于入队的栈
	outStack []int // 用于出队的栈
}

func Constructor() CQueue {
	return CQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

func (q *CQueue) AppendTail(value int) {
	q.inStack = append(q.inStack, value)
}

func (q *CQueue) DeleteHead() int {
	if len(q.outStack) == 0 {
		if len(q.inStack) == 0 {
			return -1
		}
		// 入队栈不为空
		for len(q.inStack) > 0 {
			q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
			q.inStack = q.inStack[:len(q.inStack)-1]
		}
	}
	// 出队栈不为空
	head := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return head
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
